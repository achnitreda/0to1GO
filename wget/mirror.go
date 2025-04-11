package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"sync"

	"golang.org/x/net/html"
)

type UrlQueue struct {
	mu   sync.Mutex
	urls map[string]bool
}

// our constructor
func newUrlQueue() *UrlQueue {
	return &UrlQueue{
		urls: make(map[string]bool),
	}
}

func (q *UrlQueue) add(url string) bool {
	q.mu.Lock()
	defer q.mu.Unlock()
	if !q.urls[url] {
		q.urls[url] = true
		return true
	}
	return false
}

// get the absolute url
func getFullUrl(baseUrl, relativeUrl string) string {
	base, err := url.Parse(baseUrl)
	if err != nil {
		return ""
	}

	rel, err := url.Parse(relativeUrl)
	if err != nil {
		return ""
	}
	return base.ResolveReference(rel).String()
}

func extractCssUrls(css string) []string {
	urlRegex := regexp.MustCompile(`url\(['"]?([^'"()]+)['"]?\)`)
	matches := urlRegex.FindAllStringSubmatch(css, -1)

	var urls []string
	for _, match := range matches {
		if len(match) > 1 {
			urls = append(urls, match[1])
		}
	}
	return urls
}

func convertLinksinHtml(doc *html.Node, baseUrl string) {
	var convertNode func(n *html.Node)
	convertNode = func(n *html.Node) {
		var attr string
		if n.Type == html.ElementNode {
			for i, a := range n.Attr {
				if a.Key == "style" {
					urls := extractCssUrls(a.Val)
					for _, cssUrl := range urls {
						if newUrl := getFullUrl(baseUrl, cssUrl); newUrl != "" {
							relUrl, _ := filepath.Rel(baseUrl, newUrl)
							n.Attr[i].Val = strings.Replace(
								n.Attr[i].Val,
								cssUrl,
								relUrl,
								1,
							)
						}
					}
				}
			}

			switch n.Data {
			case "a", "link":
				attr = "href"
			case "img":
				attr = "src"
			case "style":
				if n.FirstChild != nil && n.FirstChild.Type == html.TextNode {
					cssContent := n.FirstChild.Data
					urls := extractCssUrls(n.FirstChild.Data)
					for _, cssUrl := range urls {
						if newUrl := getFullUrl(baseUrl, cssUrl); newUrl != "" {
							relUrl, _ := filepath.Rel(baseUrl, newUrl)
							cssContent = strings.Replace(cssContent, cssUrl, relUrl, 1)
						}
					}
					n.FirstChild.Data = cssContent
				}
			}
		}
		if attr != "" {
			for i, a := range n.Attr {
				if a.Key == attr {
					if newUrl := getFullUrl(baseUrl, a.Val); newUrl != "" {
						relUrl, _ := filepath.Rel(baseUrl, newUrl)
						n.Attr[i].Val = relUrl
					}
					break
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			convertNode(c)
		}
	}
	convertNode(doc)
}

func MirrorWebsite(config *Config) error {
	var baseUrl *url.URL
	var err error

	// 1
	if IsValidUrl(config.url) {
		if !strings.Contains(config.url, "://") {
			config.url = "https://" + config.url
		}
		baseUrl, err = url.Parse(config.url)
		if err != nil || (baseUrl.Scheme != "http" && baseUrl.Scheme != "https") {
			return fmt.Errorf("invalid URL")
		}
	} else {
		return fmt.Errorf("invalid URL")
	}

	// 2
	outputDir := filepath.Join(".", baseUrl.Hostname())
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		return fmt.Errorf("failed to create output directory: %v", err)
	}

	// 3
	queue := newUrlQueue()
	queue.add(config.url)

	// reject & exclude
	rejectPatterns := strings.Split(config.reject, ",")
	excludeDirs := strings.Split(config.exclude, ",")

	// 4
	var wg sync.WaitGroup
	maxConcurrent := 5
	sem := make(chan struct{}, maxConcurrent)

	var processUrl func(url string) error
	processUrl = func(url string) error {
		sem <- struct{}{}
		defer func() {
			<-sem
		}()

		for _, dir := range excludeDirs {
			if dir != "" && strings.Contains(url, dir) {
				return nil
			}
		}

		client := &http.Client{}
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return fmt.Errorf("error creating request: %v", err)
		}

		// Set User-Agent to mimic Wget
		req.Header.Set("User-Agent", "Wget/1.21.2")

		resp, err := client.Do(req)
		if err != nil {
			return fmt.Errorf("failed to download %s: %v", url, err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			return fmt.Errorf("failed to download %s: %s", url, resp.Status)
		}

		// relative path
		relPath, err := filepath.Rel(baseUrl.String(), url)
		if err != nil {
			relPath = filepath.Base(url)
		}

		for _, pattern := range rejectPatterns {
			if pattern != "" && strings.HasSuffix(relPath, pattern) {
				return nil
			}
		}

		fullPath := filepath.Join(outputDir, relPath)
		if strings.HasSuffix(fullPath, "/") || filepath.Base(fullPath) == filepath.Base(outputDir) {
			fullPath = filepath.Join(fullPath, "index.html")
		}

		if err := os.MkdirAll(filepath.Dir(fullPath), 0755); err != nil {
			return fmt.Errorf("failed to create directory: %v", err)
		}

		contentType := resp.Header.Get("content-type")
		if strings.Contains(contentType, "text/html") {
			doc, err := html.Parse(resp.Body)
			if err != nil {
				return fmt.Errorf("failed to parse HTML: %v", err)
			}

			// convert-links
			if config.convertLinks {
				convertLinksinHtml(doc, baseUrl.String())
			}

			file, err := os.Create(fullPath)
			if err != nil {
				return fmt.Errorf("failed to create file: %v", err)
			}
			defer file.Close()

			if err := html.Render(file, doc); err != nil {
				return fmt.Errorf("failed to save HTML: %v", err)
			}

			var processNode func(n *html.Node)
			processNode = func(n *html.Node) {
				if n.Type == html.ElementNode {

					for _, a := range n.Attr {
						if a.Key == "style" {
							urls := extractCssUrls(a.Val)
							for _, cssUrl := range urls {
								if newUrl := getFullUrl(baseUrl.String(), cssUrl); newUrl != "" {
									if strings.HasPrefix(newUrl, baseUrl.String()) {
										if queue.add(newUrl) {
											wg.Add(1)
											go func() {
												defer wg.Done()
												processUrl(newUrl)
											}()
										}
									}
								}
							}
						}
					}

					var attr string
					/*
						-> TODO:
						// The tags that will be used for this retrieval
						// must be a, link and img that contains attributes href and src.
						-> I think no need for script to be added...
					*/
					switch n.Data {
					case "a", "link":
						attr = "href"
					case "img":
						attr = "src"
					case "style":
						if n.FirstChild != nil && n.FirstChild.Type == html.TextNode {
							urls := extractCssUrls(n.FirstChild.Data)
							for _, cssUrl := range urls {
								if newUrl := getFullUrl(baseUrl.String(), cssUrl); newUrl != "" {
									if strings.HasPrefix(newUrl, baseUrl.String()) {
										if queue.add(newUrl) {
											wg.Add(1)
											go func() {
												defer wg.Done()
												processUrl(newUrl)
											}()
										}
									}
								}
							}
						}
					}
					if attr != "" {
						for _, a := range n.Attr {
							if a.Key == attr {
								newUrl := getFullUrl(url, a.Val)
								if newUrl != "" && strings.HasPrefix(newUrl, baseUrl.String()) {
									if queue.add(newUrl) {
										wg.Add(1)
										go func() {
											defer wg.Done()
											processUrl(newUrl)
										}()
									}
								}
								break
							}
						}
					}
				}
				for c := n.FirstChild; c != nil; c = c.NextSibling {
					processNode(c)
				}
			}
			processNode(doc)
		} else {
			// non html files
			file, err := os.Create(fullPath)
			if err != nil {
				return fmt.Errorf("failed to create file: %v", err)
			}
			defer file.Close()

			if _, err := io.Copy(file, resp.Body); err != nil {
				return fmt.Errorf("failed to save file: %v", err)
			}
		}

		fmt.Printf("Downloaded: %s\n", url)
		return nil
	}

	if err := processUrl(config.url); err != nil {
		return err
	}

	wg.Wait()
	return nil
}

package main

import (
	"fmt"
	"os"
	"path"
	"strings"
	"sync"
)

func ReadUrlsFromFile(inputFile string) ([]string, error) {
	content, err := os.ReadFile(inputFile)
	if err != nil {
		return nil, err
	}
	var urls []string
	for _, line := range strings.Split(string(content), "\n") {
		if url := strings.TrimSpace(line); url != "" {
			urls = append(urls, url)
		}
	}
	return urls, nil
}

func DownloadFromFile(urls []string, baseConfig *Config) {
	var wg sync.WaitGroup
	var downloadedUrls []string
	var validUrls []string

	for _, url := range urls {
		if IsValidUrl(url) {
			if !strings.Contains(url, "://") {
				url = "https://" + url
			}
			validUrls = append(validUrls, url)
		}
	}

	results := make(chan string, len(validUrls))

	for _, url := range validUrls {
		wg.Add(1)

		go func(url string) {
			defer wg.Done()

			// note: here exactly we make copy from our original config
			// becasuse each we have each config for specific download
			config := *baseConfig
			config.url = url

			if err := downloadFile(&config); err != nil {
				results <- fmt.Sprintf("Error downloading %s: %v", url, err)
			} else {
				downloadedUrls = append(downloadedUrls, path.Base(url))
				results <- fmt.Sprintf("finished %s", path.Base(url))
			}
		}(url)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	for result := range results {
		fmt.Println(result)
	}

	fmt.Printf("\nDownload finished: %v\n", downloadedUrls)
}

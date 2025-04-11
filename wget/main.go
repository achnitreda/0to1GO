package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"
)

type Config struct {
	url          string
	outputFile   string
	outputPath   string
	rateLimit    string
	backgound    bool
	inputFile    string
	mirror       bool
	convertLinks bool
	reject       string
	exclude      string
}

func main() {
	config, err := ParseFlags()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
	startTime := GetFormattedTime()

	if config.mirror {
		fmt.Printf("start at %s\n", startTime)
		if err := MirrorWebsite(config); err != nil {
			fmt.Printf("Error mirroring website: %v\n", err)
			os.Exit(1)
		}
	} else {
		// Backgound downloading case
		if config.backgound {
			fmt.Println(`Output will be written to "wget-log"`)
			logFile, err := os.Create("wget-log")
			if err != nil {
				fmt.Printf("Error creating log file: %v\n", err)
				os.Exit(1)
			}
			defer logFile.Close()

			oldStdout := os.Stdout
			os.Stdout = logFile
			defer func() {
				os.Stdout = oldStdout
			}()

			fmt.Printf("start at %s\n", startTime)

			if err := downloadFile(config); err != nil {
				fmt.Printf("Error downloading file: %v\n", err)
			}
			endTime := GetFormattedTime()
			fmt.Printf("finished at %s\n", endTime)

			return
		}

		// handle -i flag
		if config.inputFile != "" {
			urls, err := ReadUrlsFromFile(config.inputFile)
			if err != nil {
				fmt.Printf("Error reading URLs: %v\n", err)
				os.Exit(1)
			}

			fmt.Printf("start at %s\n", startTime)
			DownloadFromFile(urls, config)
			endTime := GetFormattedTime()
			fmt.Printf("finished at %s\n", endTime)
			return
		}

		// Normal dowloading case
		fmt.Printf("start at %s\n", startTime)
		if err := downloadFile(config); err != nil {
			fmt.Printf("Error downloading file: %v\n", err)
			os.Exit(1)
		}
	}

	endTime := GetFormattedTime()
	fmt.Printf("finished at %s\n", endTime)
}

func downloadFile(config *Config) error {
	client := &http.Client{}
	req, err := http.NewRequest("GET", config.url, nil)
	if err != nil {
		return fmt.Errorf("error creating request: %v", err)
	}

	// Set User-Agent to mimic Wget
	/*
	wget -V                                                                                                                                                                                        ─╯
	GNU Wget 1.21.2 built on linux-gnu.
	*/
	req.Header.Set("User-Agent", "Wget/1.21.2")

	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("error making request: %v", err)
	}
	defer resp.Body.Close()

	fmt.Printf("sending request, awaiting response... status %s\n", resp.Status)

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status: %s", resp.Status)
	}

	contentLength := resp.ContentLength
	if contentLength > 0 {
		fmt.Printf("content size: %s\n", FormatContentLength(contentLength))
	} else {
		fmt.Println("content size: unknown")
	}

	fileName := config.outputFile
	if fileName == "" {
		fileName = path.Base(config.url)
		/*
			-> when hostname is https://theuselessweb.com
			-> then baseurl is theuselessweb.com
			-> which means that fileName is baseUrl
			-> in case basename is the filename makes it index.html
			-> like wget behaviour
		*/
		baseUrl, _ := url.Parse(config.url)
		if fileName == baseUrl.Hostname() {
			fileName = "index.html"
		}
	}

	outputPath := "."
	if config.outputPath != "" {
		outputPath = config.outputPath
		if strings.HasPrefix(outputPath, "~/") {
			home, err := os.UserHomeDir()
			if err == nil {
				outputPath = filepath.Join(home, outputPath[2:])
			}
		}
		if err := os.MkdirAll(outputPath, 0755); err != nil {
			return fmt.Errorf("error creating directory: %v", err)
		}
	}

	fullPath := filepath.Join(outputPath, fileName)
	fmt.Printf("saving file to: %s\n", fullPath)

	file, err := os.OpenFile(fullPath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return fmt.Errorf("error creating file: %v", err)
	}
	defer file.Close()

	reader := resp.Body
	if config.rateLimit != "" {
		reader, err = NewRateLimitReader(reader, config.rateLimit)
		if err != nil {
			return fmt.Errorf("error setting rate limit: %v", err)
		}
	}

	/*--- progress bar---
	-> to take a look over a buffer size tradeoffs
	-> 32KB is just recommanded from AI
	-> as said it's a balanced size that is fast and do not cost a lot of memory usage
	-> to be discussed
	*/
	buffer := make([]byte, 32*1024)
	downloaded := int64(0)
	startTime := time.Now()

	for {
		n, err := reader.Read(buffer)
		if err != nil && err != io.EOF {
			return err
		}

		if n == 0 {
			if err == io.EOF {
				break
			}
			continue
		}

		if _, err := file.Write(buffer[:n]); err != nil {
			return err
		}

		downloaded += int64(n)

		elapsed := time.Since(startTime).Seconds()
		speed := float64(downloaded) / elapsed / (1024 * 1024)

		if contentLength > 0 {
			percent := float64(downloaded) / float64(contentLength) * 100
			width := 50
			completed := int(float64(width) * float64(downloaded) / float64(contentLength))
			bar := strings.Repeat("=", completed) + strings.Repeat(" ", width-completed)

			remainingSecs := "unknown"
			if speed > 0 {
				remaining := float64(contentLength-downloaded) / (speed * 1024 * 1024)
				remainingSecs = fmt.Sprintf("%ds", int(remaining))
			}
			if config.inputFile == "" && !config.backgound {
				/*
					-> \033[K clears from cursor to end of line
					-> nsures no old characters remain visible
					-> glitch problem
				*/
				fmt.Printf("\r\033[K%s / %s [%s] %.2f%% %.2f MiB/s %s",
					FormatSize(downloaded),
					FormatSize(contentLength),
					bar,
					percent,
					speed,
					remainingSecs)
			}
		} else {
			if !config.backgound {
				fmt.Printf("\r\033[K%s downloaded %.2f MiB/s",
					FormatSize(downloaded),
					speed)
			}
		}
	}

	if !config.backgound {
		fmt.Printf("\n\nDownloaded [%s]\n", config.url)
	} else {
		fmt.Printf("Downloaded [%s]\n", config.url)
	}
	return nil
}

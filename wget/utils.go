package main

import (
	"fmt"
	"math"
	"net/url"
	"strings"
	"time"
)

func GetFormattedTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func FormatContentLength(bytes int64) string {
	const (
		MB float64 = 1024 * 1024
		GB         = MB * 1024
	)

	size := float64(bytes)

	/*
		TODO: check rounded to Mb or Gb
		-> res: content size: 56370 [~0.05MB]
		-> expected: ... [~0.06MB]
		-> Multiply by 100 before ceiling and divide by 100 after to maintain 2 decimal places
		=> it should be rounded up
	*/
	if size >= GB {
		gbSize := math.Ceil(size/GB*100) / 100
		return fmt.Sprintf("%d [~%.2fGB]", bytes, gbSize)
	} else {
		mbSize := math.Ceil(size/MB*100) / 100
		return fmt.Sprintf("%d [~%.2fMB]", bytes, mbSize)
	}
}

func FormatSize(bytes int64) string {
	const (
		KB float64 = 1024
		MB         = KB * 1024
	)

	size := float64(bytes)

	if size >= MB {
		return fmt.Sprintf("%.2f MiB", size/MB)
	} else {
		return fmt.Sprintf("%.2f KiB", size/KB)
	}
}

func IsValidUrl(arg string) bool {
	schemes := []string{"http://", "https://", "ftp://", "sftp://"}

	for _, scheme := range schemes {
		if strings.HasPrefix(arg, scheme) {
			if url, err := url.Parse(arg); err == nil && url.Host != "" {
				return true
			}
		}
	}

	// pbs.twimg.com/media/EMtmPFLWkAA8CIS.jpg
	if strings.Contains(arg, ".") && !strings.Contains(arg, " ") {
		if url, err := url.Parse("https://" + arg); err == nil && url.Host != "" {
			return true
		}
	}

	return false
}

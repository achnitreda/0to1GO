package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func ParseFlags() (*Config, error) {
	config := &Config{}

	fs := flag.NewFlagSet("wget", flag.ContinueOnError)

	fs.StringVar(&config.outputFile, "O", "", "output filename")
	fs.StringVar(&config.outputPath, "P", "", "Path to save the file")
	fs.StringVar(&config.rateLimit, "rate-limit", "", "Limit download speed (e.g., 400k for 400 KB/s, 2M for 2 MB/s)")
	fs.BoolVar(&config.backgound, "B", false, "Download file in the background")
	fs.StringVar(&config.inputFile, "i", "", "File containing URLs to download")

	// mirror flags
	fs.BoolVar(&config.mirror, "mirror", false, "Mirror a website")
	fs.BoolVar(&config.convertLinks, "convert-links", false, "Convert links for offline viewing")

	fs.StringVar(&config.reject, "reject", "", "Reject file patterns (comma-separated)")
	fs.StringVar(&config.exclude, "exclude", "", "Exclude directories (comma-separated)")

	var rejectShort string
	var excludeShort string
	fs.StringVar(&rejectShort, "R", "", "Reject file patterns (comma-separated)")
	fs.StringVar(&excludeShort, "X", "", "Exclude directories (comma-separated)")

	var url string
	args := os.Args[1:]
	reorganizedArgs := []string{}

	for i := 0; i < len(args); i++ {
		arg := args[i]
		if strings.HasPrefix(arg, "-") {
			reorganizedArgs = append(reorganizedArgs, arg)
			if arg != "--convert-links" && arg != "--mirror" && arg != "-B" && !strings.Contains(arg, "=") && i+1 < len(args) && !strings.HasPrefix(args[i+1], "-") {
				reorganizedArgs = append(reorganizedArgs, args[i+1])
				i++
			}
		} else if url == "" && IsValidUrl(arg) {
			url = arg
			if !strings.Contains(arg, "://") {
				url = "https://" + arg
			}
		}
	}

	if err := fs.Parse(reorganizedArgs); err != nil {
		return nil, err
	}

	if rejectShort != "" {
		config.reject = rejectShort
	}

	if excludeShort != "" {
		config.exclude = excludeShort
	}

	// If -i flag is present, don't require URL
	if config.inputFile != "" {
		return config, nil
	}

	if url == "" {
		return nil, fmt.Errorf("no valid URL found in arguments")
	}
	config.url = url
	return config, nil
}

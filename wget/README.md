# Go wget Implementation

This project implements a simplified version of the GNU `wget` utility in Go. It provides functionality for downloading files from the web with various options and features.

## Features

- Download a single file from a URL
- Save files under different names
- Save files to specific directories
- Limit download speed
- Download files in the background
- Download multiple files concurrently
- Mirror entire websites

## Usage

### Basic Usage

Download a file from a URL:

```bash
go run . https://example.com/file.zip
```

### Command-Line Options

#### Change Output Filename (`-O`)

```bash
go run . -O=myfile.jpg https://example.com/image.jpg
```

#### Specify Output Directory (`-P`)

```bash
go run . -P=~/Downloads/ https://example.com/file.zip
```

#### Limit Download Speed (`--rate-limit`)

Limit downloads to specific speeds (k = KB/s, M = MB/s):

```bash
go run . --rate-limit=300k https://example.com/bigfile.zip
go run . --rate-limit=2M https://example.com/hugefile.zip
```

#### Background Download (`-B`)

Download a file in the background and write output to a log file:

```bash
go run . -B https://example.com/file.zip
```

#### Download Multiple Files (`-i`)

Download multiple URLs listed in a file:

```bash
go run . -i=urls.txt
```

Where `urls.txt` contains a list of URLs, one per line.

#### Mirror a Website (`--mirror`)

Download an entire website for offline viewing:

```bash
go run . --mirror https://example.com
```

#### Mirror with Link Conversion (`--convert-links`)

Download a website and convert links for offline viewing:

```bash
go run . --mirror --convert-links https://example.com
```

#### Reject File Types (`-R` or `--reject`)

Mirror a website but exclude specific file types:

```bash
go run . --mirror -R=jpg,gif,png https://example.com
```

#### Exclude Directories (`-X`)

Mirror a website but exclude specific directories:

```bash
go run . --mirror -X=/images,/css https://example.com
```

## Output Format

The program provides detailed feedback during downloads:

```
start at 2024-04-11 15:04:05
sending request, awaiting response... status 200 OK
content size: 56370 [~0.06MB]
saving file to: ./myfile.jpg
 55.05 KiB / 55.05 KiB [================================================] 100.00% 1.24 MiB/s 0s
Downloaded [https://example.com/image.jpg]
finished at 2024-04-11 15:04:07
```

## Dependencies

This project uses only the Go standard library, with no external dependencies.

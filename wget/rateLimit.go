package main

import (
	"fmt"
	"io"
	"strconv"
	"strings"
	"time"
)

type RateLimitReader struct {
	r            io.ReadCloser
	rateLimit    int64
	lastTimeRead time.Time
	bytesRead    int64
}

// Read implements io.ReadCloser.
func (r *RateLimitReader) Read(p []byte) (n int, err error) {
	now := time.Now()

	if r.lastTimeRead.IsZero() {
		n, err = r.r.Read(p)
		r.bytesRead += int64(n)
		r.lastTimeRead = now
		return
	}

	elapsed := now.Sub(r.lastTimeRead).Seconds()

	allowedBytes := int64(elapsed * float64(r.rateLimit))

	if len(p) > int(allowedBytes) {
		p = p[:allowedBytes]
		if allowedBytes == 0 {
			time.Sleep(time.Millisecond * 100)
			return 0, nil
		}
	}
	n, err = r.r.Read(p)
	r.bytesRead += int64(n)
	r.lastTimeRead = now

	return
}

// Close implements io.ReadCloser.
func (r *RateLimitReader) Close() error {
	return r.r.Close()
}

func NewRateLimitReader(r io.ReadCloser, rateLimit string) (io.ReadCloser, error) {
	var bytesPerSec int64

	if rateLimit == "" {
		return r, nil
	}

	value := rateLimit[:len(rateLimit)-1]
	unit := strings.ToLower(rateLimit[len(rateLimit)-1:])

	var multiplier int64
	switch unit {
	case "k":
		multiplier = 1024
	case "m":
		multiplier = 1024 * 1024
	default:
		return nil, fmt.Errorf("invalid rate limit unit: %s", unit)
	}

	val, err := strconv.Atoi(value)
	if err != nil {
		return nil, fmt.Errorf("invalid rate limit value: %s", value)
	}

	bytesPerSec = int64(val) * multiplier

	return &RateLimitReader{
		r:            r,
		rateLimit:    bytesPerSec,
		lastTimeRead: time.Time{},
	}, nil
}

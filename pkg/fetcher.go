package pkg

import (
	"fmt"
	"net/http"
	"time"
)

type Result struct {
	URL      string
	Code     int
	Size     int64
	Duration time.Duration
	Error    error
}

func Fetch(url string) Result {
	start := time.Now()

	resp, err := http.Get(url)
	if err != nil {
		return Result{URL: url, Error: fmt.Errorf("\033[1;33mERROR:\033[0m %v", err)}
	}
	defer resp.Body.Close()

	duration := time.Since(start)
	return Result{URL: url, Code: resp.StatusCode, Size: resp.ContentLength, Duration: duration}
}

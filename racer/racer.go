package racer

import (
	"net/http"
	"time"
)

func Racer(a string, b string) (winner string) {
	aTime := measureResponseTime(a)
	bTime := measureResponseTime(b)

	if aTime < bTime {
		return a
	}

	return b
}

func measureResponseTime(url string) time.Duration {
	startTime := time.Now()
	http.Get(url)
	return time.Since(startTime)
}

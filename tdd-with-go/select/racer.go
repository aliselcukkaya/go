package racer

import (
	"net/http"
	"time"
)

func Racer(firstURL, secondURL string) (winner string) {
	firstURLResponseTime := measureResponseTime(firstURL)
	secondURLResponseTime := measureResponseTime(secondURL)

	if firstURLResponseTime < secondURLResponseTime {
		return firstURL
	}
	return secondURL
}

func measureResponseTime(url string) time.Duration {
	startURL := time.Now()
	http.Get(url)
	return time.Since(startURL)
}

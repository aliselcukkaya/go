package racer

import (
	"fmt"
	"net/http"
	"time"
)

var tenSecondTimer = 10 * time.Second

func Racer(firstURL, secondURL string) (winner string, error error) {
	return ConfigurableRacer(firstURL, secondURL, tenSecondTimer)
}

func ConfigurableRacer(firstURL, secondURL string, timeout time.Duration) (winner string, error error) {
	select {
	case <-ping(firstURL):
		return firstURL, nil
	case <-ping(secondURL):
		return secondURL, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("timed out waiting for %s and %s", firstURL, secondURL)
	}
}

func ping(url string) chan struct{} {
	ch := make(chan struct{})

	go func() {
		http.Get(url)
		close(ch)
	}()
	return ch
}

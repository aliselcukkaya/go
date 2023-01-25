package racer

import "net/http"

func Racer(firstURL, secondURL string) (winner string) {
	select {
	case <-ping(firstURL):
		return firstURL
	case <-ping(secondURL):
		return secondURL
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

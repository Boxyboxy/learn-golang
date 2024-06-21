package sel

import (
	"fmt"
	"net/http"
	"time"
)

/*
	func Racer(a, b string) (winner string) {
		aDuration := measureReponseTime(a)

		bDuration := measureReponseTime(b)

		if aDuration < bDuration {
			return a
		}

		return b
	}

func measureReponseTime(url string) time.Duration {
start := time.Now()

		http.Get(url)
		return time.Since(start)
	}
*/
var tenSecondTimeout = 10 * time.Second

func Racer(a, b string) (winner string, error error) {
	return ConfigurableRacer(a, b, tenSecondTimeout)
}
func ConfigurableRacer(a, b string, timeout time.Duration) (winner string, error error) {
	select {
	// no received values, the following syntax just waits for the channels to be closed
	// select is just used to determine which requests finishes first
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("timed out waiting for %s and %s", a, b)
	}
}

// In the following code, nothing is being communicated through the channel
func ping(url string) chan struct{} {
	ch := make(chan struct{})
	go func() {
		http.Get(url)
		close(ch)
	}()
	return ch
}

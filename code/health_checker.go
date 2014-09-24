// +build OMIT

package main

import (
	"fmt"
	"net/http"
)

// START OMIT
func checkHTTPStatus(url string, out chan string) {
	resp, _ := http.Get(url)
	out <- fmt.Sprintf("%s : [ %s ]", url, resp.Status)
}

func main() {
	statusCodes := make(chan string)

	var urls = []string{
		"http://www.mozilla.org", "http://www.coursera.org", "http://www.eff.org",
		"http://www.wc3.org", "http://www.torproject.org", "http://www.repubblica.it",
		"http://www.ubuntu.com/404", "http://127.0.0.1:3999",
	}

	for _, url := range urls {
		go checkHTTPStatus(url, statusCodes)
	}

	for _ = range urls {
		fmt.Println(<-statusCodes)
	}
}

// END OMIT

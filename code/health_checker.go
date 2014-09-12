// +build OMIT

package main

import (
	"fmt"
	"net/http"
)

// START OMIT
var urls = []string{
	"http://www.ubuntu.com/404", "http://www.mozilla.org", "http://www.eff.org",
}

func checkHTTPStatus(url string, out chan string) {
	resp, _ := http.Get(url)
	out <- fmt.Sprintf("%s : [ %s ]", url, resp.Status)
}

func main() {
	statusCodes := make(chan string)
	for _, url := range urls {
		go checkHTTPStatus(url, statusCodes)
	}
	for _ = range urls {
		fmt.Println(<-statusCodes)
	}
}

// END OMIT

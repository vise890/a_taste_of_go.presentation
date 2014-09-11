// +build OMIT

package main

import (
	"fmt"
	"net/http"
)

var (
	urls = []string{
		"http://www.ubuntu.com/4040404040",
		"http://www.mozilla.org",
		"http://www.example.com",
		"http://www.eff.org",
		"http://www.iclearlydontexist.com",
	}
)

// START OMIT
func checkHTTPStatus(url string, out chan string) {
	resp, err := http.Get(url)
	if err != nil {
		out <- fmt.Sprintf("%s : [ %s ]", url, err.Error())
		return
	}
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

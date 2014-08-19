// +build OMIT

package main

import (
	"fmt"
	"net/http"
)

var (
	urls = []string{
		"http://www.google.com",
		"http://www.google.com/abcdefg404",
		"http://www.ubuntu.com",
		"http://www.yoox.com",
		"http://www.example.com",
		"http://www.iclearlydontexist.com",
	}
)

// START OMIT
func getResponseCode(url string, out chan string) {
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
		go getResponseCode(url, statusCodes)
	}
	for _ = range urls {
		fmt.Println(<-statusCodes)
	}
}

// END OMIT

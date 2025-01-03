package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func checkURL(url string) {
	start := time.Now()

	// Send a GET request to the URL
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("[ERROR] %s: %s\n", url, err)
		return
	}
	defer resp.Body.Close()

	elapsed := time.Since(start)

	// Print the status code and response time
	fmt.Printf("[OK] %s: Status %d, Time %s\n", url, resp.StatusCode, elapsed)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <url1> <url2> ...")
		return
	}

	urls := os.Args[1:]
	for _, url := range urls {
		checkURL(url)
	}
}

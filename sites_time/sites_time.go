package main

import (
	"io"
	"log"
	"net/http"
	"sync"
	"time"
)

func siteTime(url string) {
	startTime := time.Now()

	response, err := http.Get(url)
	if err != nil {
		log.Printf("ERROR: %s -> %s", url, err)
		return
	}
	defer func() {
		bodyCloseErr := response.Body.Close()
		if bodyCloseErr != nil {
			log.Printf("ERROR: Closing response body for %s -> %s", url, bodyCloseErr)
		}
	}()

	if _, err := io.Copy(io.Discard, response.Body); err != nil {
		log.Printf("ERROR: Reading response body for %s -> %s", url, err)
	}

	executionTime := time.Since(startTime).Milliseconds()

	log.Printf("INFO: [%s] Execution time milliseconds: %dms", url, executionTime)
}

func main() {
	siteTime("https://google.com")

	urls := []string{
		"https://google.com",
		"https://apple.com",
		"https://ya.ru",
		"https://dzen.ru",
	}

	var waitGroup sync.WaitGroup
	for _, url := range urls {
		waitGroup.Add(1)
		go func(url string) {
			siteTime(url)
			defer waitGroup.Done()
		}(url)
	}

	waitGroup.Wait()
}

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

func crawlWeb(done <-chan struct{}, input <-chan string) <-chan string {
	output := make(chan string)

	// Create a semaphore to limit the number of concurrent goroutines
	semaphore := make(chan struct{}, 8)

	var wg sync.WaitGroup

	go func() {
		for url := range input {
			// Acquire a token from the semaphore
			semaphore <- struct{}{}
			time.Sleep(time.Second)                        //// SLEEP
			// Launch a goroutine for each URL
			wg.Add(1)
			go func(url string) {
				defer func() {
					// Release the token back to the semaphore
					<-semaphore
					wg.Done()
				}()

				// Send a GET request to the URL
				resp, err := http.Get(url)
				if err != nil {
					fmt.Println(err)
					return
				}
				defer resp.Body.Close()

				// Read the response body into a string
				body, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					fmt.Println(err)
					return
				}

				// Send the response body to the output channel
				// output <- string(body)
				select {
				case output <- string(body):
				case <-done:
					return
				}
			}(url)
		}

		// Wait for all the goroutines to finish
		wg.Wait()

		// Close the output channel when all URLs have been processed
		close(output)
	}()

	return output
}

func main() {
	done := make(chan struct{})
	input := make(chan string)

	// Start the crawlWeb function
	output := crawlWeb(done, input)

	// Send some URLs to the input channel
	go func() {
		for _, url := range []string{
			// "https://www.google.com",
			// "https://www.bing.com",
			// "https://www.yahoo.com",
			// "https://www.duckduckgo.com",
			// "https://repos.21-school.ru/users/sign_in",
			"http://example.com",
		} {
			input <- url
		}
		close(input)
	}()

	// Read the results from the output channel
	for body := range output {
		fmt.Println(body)
		fmt.Printf("\nnew:\n")
	}
	close(done)
}

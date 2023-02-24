package main

import (
	"fmt"
	"net/http"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(200)
	for i := 0; i < 200; i++ {
		go func() {
			defer wg.Done()
			res, err := http.Get("http://localhost:8888/")
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
			defer res.Body.Close()
			fmt.Println("Response status:", res.Status)
		}()
	}
	wg.Wait()
}

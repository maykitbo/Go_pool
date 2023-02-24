package main

import (
	"fmt"
	"sync"
	"math/rand"
	"time"
)

func SleepSort(arr []int) {
	var wg sync.WaitGroup
	k := len(arr)
	wg.Add(k)
	for t := 0; t < k; t++ {
		go func(n int, wg *sync.WaitGroup) {
			defer wg.Done()
			time.Sleep(time.Duration(n) * time.Second)
			fmt.Println(n)
		}(arr[t], &wg)
	}
	wg.Wait()
}

func main() {
	t := rand.Perm(15)
	fmt.Println(t)
	SleepSort(t)
	SleepSort([]int{3, 5, 7})
}

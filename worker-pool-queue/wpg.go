package main

import (
	"fmt"
	"sync"
	"time"
)

var concurrency = 5

func main() {
	// put tasks on channel
	tasks := make(chan int, 100)
	go func() {
		for j := 1; j <= 9; j++ {
			tasks <- j
		}
		close(tasks)
	}()

	// waitgroup, and close results channel when work done
	results := make(chan string)
	wg := &sync.WaitGroup{}
	wg.Add(concurrency)
	go func() {
		wg.Wait()
		close(results)
	}()

	for i := 1; i <= concurrency; i++ {
		go func(id int) {
			defer wg.Done()

			for t := range tasks {
				fmt.Println("worker", id, "processing job", t)
				results <- string(t * 2)
				time.Sleep(time.Second)
			}
		}(i)
	}

	// loop over results until closed (see above)
	for r := range results {
		fmt.Printf("result %v\n", r)
	}
}

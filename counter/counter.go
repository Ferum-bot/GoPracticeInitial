package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	count := int64(0)

	var wait sync.WaitGroup
	for i := 0; i < 20; i++ {
		wait.Add(1)
		go func() {
			for j := 0; j < 1_000; j++ {
				atomic.AddInt64(&count, 1)
			}
			defer wait.Done()
		}()
	}

	wait.Wait()
	fmt.Println("Count: ", count)
}

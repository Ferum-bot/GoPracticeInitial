package main

import (
	"fmt"
	"time"
)

/*
	Channel semantics

- send & receive will block until opposite operation (*)
- receive from a closed channel will return the zero value without blocking
- send to empty channel raise panic
- closing a closed channel raise panic
- send/receive to nil channel will block forever
*/
func main() {
	go fmt.Println("Hello from side goroutine")
	fmt.Printf("Hello from main goroutine\n")

	for i := 0; i < 3; i++ {
		go func(localI int) {
			fmt.Printf("Hello from goroutine number %d\n", localI)
		}(i)
	}

	time.Sleep(20 * time.Millisecond)

	channel := make(chan string)
	go func() {
		channel <- "Hello from main goroutine"
	}()
	messageFromChannel := <-channel
	fmt.Println(messageFromChannel)

	go func() {
		for i := 0; i < 3; i++ {
			message := fmt.Sprintf("Message number %d", i+1)
			channel <- message
		}
		close(channel)
	}()

	for message := range channel {
		fmt.Println("Received: ", message)
	}

	emptyMessage, ok := <-channel
	fmt.Printf("Message from closed channel: %#v, is ok: %#v\n", emptyMessage, ok)

	fmt.Println(sleepSort([]int{9, 8, 7, 6, 5, 4, 3, 2, 1}))
}

func sleepSort(values []int) []int {
	channel := make(chan int)
	for _, value := range values {
		go sleepAndProduce(value, channel)
	}

	sortedValues := make([]int, 0, len(values))
	for value := range channel {
		sortedValues = append(sortedValues, value)

		if len(sortedValues) == len(values) {
			close(channel)
		}
	}

	return sortedValues
}

func sleepAndProduce(value int, channel chan int) {
	sleepTime := time.Duration(value) * time.Millisecond
	time.Sleep(sleepTime)
	channel <- value
}

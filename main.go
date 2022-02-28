package main

import (
	"fmt"
	"queue/queue"
	"sync"
)

func main() {
	fmt.Println("Hello New World")
	myQueue := queue.NewQueue()
	my_channel := make(chan int)

	func() {
		for i :=0; i<1000; i++ {
			myQueue.Add(i)
		}
	}()

	go func() {
		wg := sync.WaitGroup{}
		for i :=0; i<1000; i++ {
			wg.Add(1)
			a, _ := myQueue.Remove()
			my_channel <- a
			wg.Done()
		}
		wg.Wait()
		close(my_channel)
	}()
		
	for n := range my_channel {		
		fmt.Printf("Value of n : %d\n", n)
	}

}
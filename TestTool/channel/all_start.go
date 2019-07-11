package main

import (
	"fmt"
	"sync"
)

// 通过channel阻塞goroutine， close(chan)以同时唤醒goroutine
func main() {
	begin := make(chan interface{})
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			<-begin
			fmt.Printf("%v has begun\n", i)
		} (i)
	}

	fmt.Println("Unblocking goroutines...")
	close(begin)
	wg.Wait()
}

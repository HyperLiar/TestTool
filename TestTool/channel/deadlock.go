package main

import "fmt"

func main() {
	// fatal error: all goroutines are asleep - deadlock!
	//var dataChan chan interface{}
	//<-dataChan

	// goroutine阻塞
	/*go func() {
		var dataChan chan interface{}
		<-dataChan
	} ()*/

	// fatal error: all goroutines are asleep - deadlock!
	/*var dataChan chan interface{}
	dataChan <- struct{}{}*/

	chanOwner := func() <-chan int {
		resultChan := make(chan int, 5)
		go func() {
			defer close(resultChan)
			for i := 0; i <= 5; i++ {
				resultChan <- i
			}
		} ()

		return resultChan
	}

	resultChan := chanOwner()
	for result := range resultChan {
		fmt.Printf("Received: %d\n", result)
	}

	fmt.Println("Done receiving!")
}

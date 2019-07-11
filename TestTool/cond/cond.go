package main

import (
	"fmt"
	"sync"
	"time"
)

var m sync.Mutex

// 可以使用channel
func main() {
	c := sync.NewCond(&m)

	//ready := make(chan struct{})
	isReady := false

	for i := 0; i < 10; i++ {
		a := i
		go func() {
			m.Lock()
			// time.Sleep(10*time.Millisecond)
			//ready <- struct{}{}
			for !isReady {
				c.Wait()
			}

			fmt.Println(a)
			m.Unlock()
		} ()
	}

	c.Broadcast()
	// c.Signal()
	time.Sleep(1*time.Second)

	for i := 0; i < 10; i++ {
		//<-ready
	}
	isReady = true
	c.Broadcast()
	time.Sleep(1*time.Second)
}

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// leak1()
	leak2Fix()
}

func leak1() {
	doWork := func(strings <-chan string) <-chan interface{} {
		completed := make(chan interface{})
		go func() {
			defer fmt.Println("doWork exited.")
			defer close(completed)

			for s := range strings {
				fmt.Println(s)
			}
		}()

		return completed
	}

	// strings channel永远不会读到数据，所以goroutine永远也不会结束
	doWork(nil)
	fmt.Println("Done.")
}

func leak1Fix() {
	doWork := func(done <-chan interface{}, strings <-chan string) <-chan interface{} {
		terminated := make(chan interface{})
		go func() {
			defer fmt.Println("doWork exited.")
			defer close(terminated)

			for {
				select {
				case s := <-strings:
					fmt.Println(s)
				case t := <-done:
					// close的时候会收到信号，执行return -> defer
					fmt.Println(t)
					return
				}
			}
		}()
		return terminated
	}

	done := make(chan interface{})
	terminated := doWork(done, nil)

	go func() {
		time.Sleep(1 * time.Second)
		fmt.Println("Canceling doWork goroutine...")
		close(done)
	}()

	<-terminated
	fmt.Println("Done.")
}

func leak2() {
	newRandChan := func() <-chan int {
		randChan := make(chan int)
		go func() {
			defer fmt.Println("newRandChan closure exited.")
			defer close(randChan)

			for {
				randChan <- rand.Int()
			}
		}()
		return randChan
	}
	randChan := newRandChan()
	fmt.Println("3 random ints:")
	for i := 1; i <= 3; i++ {
		fmt.Printf("%d: %d\n", i, <-randChan)
	}
}

func leak2Fix() {
	newRandChan := func(done <-chan interface{}) <-chan int {
		randChan := make(chan int)
		go func() {
			defer fmt.Println("newRandChan closure exited.")
			defer close(randChan)

			for {
				select {
				case randChan <- rand.Int():
				case <-done:
					return
				}
			}
		}()
		return randChan
	}

	done := make(chan interface{})
	randChan := newRandChan(done)
	fmt.Println("3 random ints:")

	for i := 1; i <= 3; i++ {
		fmt.Printf("%d: %d\n", i, <-randChan)
	}
	close(done)
	//time.Sleep(1 * time.Second)
}

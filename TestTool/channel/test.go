package main

import (
	"fmt"
	"os"
	"runtime"
	"runtime/pprof"
	"time"
)

// 一直输出goroutine: 2
// 因为有1个缓冲区，22行的goroutine会直接退出
// 28行因为传入的channel为nil，会一直阻塞
// 如果去掉缓冲区，22行也会阻塞
// 如果28行goroutine去掉传值逻辑，也不会阻塞
// 本质的问题是：channel在一个goroutine内初始化，同时试图在另一个goroutine读取，且channel的初始值为nil

/**
1.第n个send一定happen before第n个receive，无论是buffered channel/unbuffered channel
2. capacity是m的channel, 第n个receive一定happen before第n+m个send
3. m = 0的unbufferd channel, 第n个receive一定happend before第n个send
4. channel的close一定happen before receive端得到通知，这意味着一个receive收到一个channel close收到的0值
 */
func main() {
	var ch chan int

	go func() {
		ch = make(chan int, 1)
		ch <- 1
	}()

	// 此时传入的channel是nil
	go func(ch chan int) {
		fmt.Println(ch)
		time.Sleep(time.Second)
		<-ch
	}(ch)

	c := time.Tick(1 * time.Second)
	counter := 0

	for range c {
		fmt.Printf("#gouroutines: %d\n", runtime.NumGoroutine())
		counter++
		if counter >= 3 {
			break
		}
	}

	pprof.Lookup("goroutine").WriteTo(os.Stdout, 1)
}

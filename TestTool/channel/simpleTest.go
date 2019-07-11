package main

import (
	"fmt"
	"time"
)

func main() {
	// empty()
	// closeRead()
	// closeWrite()
	// close2()
	// isClose()
	overTime()
}

// fatal error: all goroutines are asleep - deadlock!
func emtpy() {
	ch1 := make(chan int)

	ch1 = nil
	x := 1
	ch1 <- x
}

// 得到对应channel类型的值
func closeRead() {
	// 返回0
	ch1 := make(chan int)

	close(ch1)

	x := <-ch1
	fmt.Println(x)

	//  invalid operation: <-ch1 (receive from send-only type chan<- int)
	ch2 := make(chan<- int)

	close(ch2)

	//x := <-ch2
	//fmt.Println(x)
}

// panic: send on closed channel
func closeWrite() {
	ch := make(chan bool)
	close(ch)
	ch <- true
}

// panic: close of closed channel
func close2() {
	ch := make(chan bool)
	close(ch)
	close(ch)
}

// 尽管关闭了，仍然能读取到数据，根据第二个参数判断channel状态
func isClose() {
	ch := make(chan bool)
	i, ok := <-ch
	if ok {
		fmt.Println(i)
	} else {
		fmt.Println("channel closed")
	}
}

// 超时检查
func overTime() {
	ch := make(chan bool)
	select {
		case <-time.After(2*time.Second):
			fmt.Println("over time 2 second")
		case i := <-ch:
			fmt.Println("receive i", i)
	}
}
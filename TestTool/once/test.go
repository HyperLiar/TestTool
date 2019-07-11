package main

import (
	"fmt"
	"sync"
)

type Once struct {
	m sync.Mutex
	done uint32
}

// 并发场景会有问题
// cpu会缓存，即使done在本goroutine读取到了1，其他的goroutine可能读取到的是0
// 必须通过atomic保证happen before
// 记，获取到写锁的goroutine一定发生在读锁前
func (o *Once) Do(f func()) {
	if o.done == 1 {
		return
	}

	o.m.Lock()
	defer o.m.Unlock()

	if o.done == 0 {
		o.done = 1
		f()
	}
}

func main() {
	var o Once

	var wg sync.WaitGroup
	for i:=0; i< 10; i++ {
		wg.Add(1)
		go func() {
			o.Do(initInt)
			wg.Done()
		} ()
	}

	wg.Wait()
}

func initInt() {
	fmt.Println("call init")
}
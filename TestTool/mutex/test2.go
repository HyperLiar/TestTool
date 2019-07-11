package main

import (
	"fmt"
	"sync"
)

type MyMutex struct {
	count int
	sync.Mutex
}

// fatal error: all goroutines are asleep - deadlock!
func main() {
	var mu MyMutex

	mu.Lock()
	var mu2 = mu
	mu.count++
	mu.Unlock()

	// mu2也处在被锁住的状态，不能重复上锁
	mu2.Lock()
	mu2.count++
	mu2.Unlock()

	fmt.Println(mu.count, mu2.count)
}
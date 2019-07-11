package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var m sync.RWMutex
	rr(&m, 5)
}

// 递归读，实际是串行阻塞
func rr(m *sync.RWMutex, n int) int {
	if n < 1 {
		return 0
	}

	fmt.Println("RLock", n)
	m.RLock()
	defer func() {
		fmt.Println("RUnlock", n)
		m.RUnlock()
	}()

	time.Sleep(100 * time.Millisecond)
	return rr(m, n-1) + n
}

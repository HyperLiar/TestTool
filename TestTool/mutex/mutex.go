package main

import (
	"fmt"
	"sync"
	"time"
)

var m sync.Mutex

// 注意这里，goroutine执行并不一定是按照创建的顺序
func main() {
	var a map[int]int
	a = make(map[int]int, 5)

	a[1], a[2], a[3], a[4] = 10, 10, 10, 10

	for i := 0; i < 4; i++ {
		i := i
		go func(b map[int]int) {
			m.Lock()
			b[4] = i * 10
			fmt.Println(a)
			m.Unlock()
		}(a)
	}

	m.Lock()
	fmt.Println(a)
	m.Unlock()

	time.Sleep(time.Second)
	fmt.Println(a)
}

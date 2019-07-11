package main

import (
	"sync"
	"time"
)

var m *sync.RWMutex

// 读锁彼此不阻塞
// 写锁彼此阻塞，写锁和读锁也阻塞
// writer.Lock高于reader.RLock
func main() {
	m = new(sync.RWMutex)
	for i := 0; i < 5; i++ {
		go read(i)
	}

	for i := 0; i < 5; i++ {
		go write(i)
	}

	time.Sleep(10 * time.Second)
}

func read(i int) {
	println(i, "read start")
	m.RLock()
	var p = 0
	var pr = "read"
	for {
		pr += "."
		if p == 3 {
			break
		}
		time.Sleep(350 * time.Millisecond)
		p++
		println(i, pr)
	}
	m.RUnlock()
	println(i, "read end")
}

func write(i int) {
	println(i, "write start")
	m.Lock()

	var p = 0
	var pr = "write"
	for {
		pr += "."
		if p == 3 {
			break
		}

		time.Sleep(350 * time.Millisecond)
		p++
		println(i, pr)
	}
	m.Unlock()
	println(i, "write end")
}

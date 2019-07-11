package main

import (
	"bytes"
	"fmt"
	"net/http"
	"runtime"
	"sync"
	"time"
	_ "net/http/pprof"
)

var pool = sync.Pool{New: func() interface{} { return new(bytes.Buffer)}}

// 内存暴涨
func main() {
	go func() {
		http.ListenAndServe("localhost:6060", nil)
	} ()

	go func() {
		for {
			processRequest(1 << 20) // 256 MiB
		}
	} ()

	for i := 0; i < 100; i++ {
		go func() {
			for {
				processRequest(1 << 10) // 1KiB
			}
		} ()
	}

	var stats runtime.MemStats
	for i := 0; ; i++ {
		runtime.ReadMemStats(&stats)
		fmt.Printf("Cycle %d: %dB\n", i, stats.Alloc)
		time.Sleep(time.Second)
		runtime.GC()
	}
}

func processRequest(size int) {
	b := pool.Get().(*bytes.Buffer)
	time.Sleep(500 * time.Millisecond)
	b.Grow(size)
	pool.Put(b)
	time.Sleep(1 * time.Millisecond)
}
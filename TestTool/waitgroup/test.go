package main

import (
	"sync"
	"time"
)

// panic: sync: WaitGroup is reused before previous Wait has returned
func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		time.Sleep(time.Millisecond)
		wg.Done()
		wg.Add(1)
		// add后面必须要有对应的done, 否则wait会认为前一次wait没有处理
		// wg.Done()
	} ()
	wg.Wait()
}
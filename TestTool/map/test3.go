package main

import (
	"fmt"
	"sync"
)

// 输出可能不是2000
// 因为append并不是并发安全的，如果两个goroutine同时发现ints[1]是空的，他们都会把自己的值放在这个位置
// 两个值重合覆盖，造成数据丢失
func main() {
	var wg sync.WaitGroup

	wg.Add(2)

	var ints = make([]int, 0, 1000)

	go func() {
		for i := 0; i < 1000; i++ {
			ints = append(ints, i)
		}

		wg.Done()
	}()

	go func() {
		for i := 0; i < 1000; i ++ {
			ints = append(ints, i)
		}

		wg.Done()
	}()

	wg.Wait()

	fmt.Println(len(ints))
}

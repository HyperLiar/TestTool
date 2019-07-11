package main

import "time"

func main() {
	done := make(chan interface{})
	defer close(done)

	zeros := take(done, 3, repeat(done, 0))

	short := sleep(done, 1*time.Second, zeros)
	long := sleep(done, 4*time.Second, short)
	pipeline := long
}

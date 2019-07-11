package main

var c = make(chan int)
var a int

func f() {
	a = 1
	<- c
}

// 输出1，happens before关系。
// channel的写入一定在读取前
func main() {
	go f()
	c <- 0
	print(a)
}
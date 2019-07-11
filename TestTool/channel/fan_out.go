package main

import "reflect"

/**
	从输入channel中读取一个数据，发送给每个输入channel，这种模式称之为Tee模式
*/

// channel
func fanOut(ch <-chan interface{}, out []chan interface{}, async bool) {
	go func() {
		defer func() {
			for i := 0; i < len(out); i++ {
				close(out[i])
			}
		}()

		for v := range ch {
			v := v
			for i := 0; i < len(out); i++ {
				i := i

				if async {
					go func() {
						out[i] <- i
					}()
				} else {
					out[i] <- v
				}
			}
		}
	}()
}

// reflect
func fanOutReflect(ch <-chan interface{}, out []chan interface{}) {
	go func() {
		defer func() {
			for i := 0; i < len(out); i++ {
				close(out[i])
			}
		}()
	}()

	cases := make([]reflect.SelectCase, len(out))
	for i := range cases {
		cases[i].Dir = reflect.SelectSend
	}

	for v := range ch {
		v := v
		for i := range cases {
			cases[i].Chan = reflect.ValueOf(out[i])
			cases[i].Send = reflect.ValueOf(v)
		}

		for _ = range cases {
			chosen, _, _ := reflect.Select(cases)
			cases[chosen].Chan = reflect.ValueOf(nil)
		}
	}
}

/*
从输入channel中读取一个数据，在输出channel中选择一个channel发送
 */
func fanOut(ch <-chan interface{}, out []chan interface{}) {
	go func() {
		defer func() {
			for i := 0; i < len(out); i++ {
				close(out[i])
			}
		}()

		// roundrobin
		var i = 0
		var n = len(out)
		for v := range ch {
			v := v
			out[i] <- v
			i = (i + 1) % n
		}
	}()
}

func fanOutReflect(ch <-chan interface{}, out []chan interface{}) {
	go func() {
		defer func() {
			for i := 0; i < len(out); i++ {
				close(out[i])
			}
		}()

		cases := make([]reflect.SelectCase, len(out))

		for i := range cases {
			cases[i].Dir = reflect.SelectSend
			cases[i].Chan = reflect.ValueOf(out[i])
		}

		for v := range ch {
			v := v
			for i := range cases {
				cases[i].Send = reflect.ValueOf(v)
			}

			_, _, _ = reflect.Select(cases)
		}
	}()
}

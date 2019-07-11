package main

func reverse(x int) int {
	var max int32 = int32(^uint32(0) >> 1) / 10
	var min int32 = (-1 << 31) / 10
	var in int32 = int32(x)
	var out int32 = 0
	for ; in != 0; {
		digit := in%10
		if out < min || out > max {
			return 0
		}
		out = out*10 + digit
		in = in/10
	}
	return int(out)
}
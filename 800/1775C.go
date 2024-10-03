package main

import (
	"bufio"
	. "fmt"
	"os"
)

func main() {
	var t int
	in := bufio.NewReader(os.Stdin)
	for Fscan(in, &t); t > 0; t-- {
		var n, x int64
		Fscan(in, &n, &x)
		var res int64 = n

		for res > x {
			n += n & -n
			res &= n
		}

		if res == x {
			Println(n)
		} else {
			Println(-1)
		}
	}
	return
}

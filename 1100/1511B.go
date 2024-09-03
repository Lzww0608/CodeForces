package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var t, a, b, c int
	for fmt.Fscan(in, &t); t > 0; t-- {
		fmt.Fscan(in, &a, &b, &c)
		x := make([]byte, a)
		y := make([]byte, b)
		x[0], y[0] = '1', '1'
		for i := 1; i < a; i++ {
			x[i] = '0'
		}
		for i := 1; i < b; i++ {
			y[i] = '0'
		}
		if a > b {
			x[a-c] = '1'
		} else {
			y[b-c] = '1'
		}
		fmt.Fprintln(out, string(x), string(y))
	}
}

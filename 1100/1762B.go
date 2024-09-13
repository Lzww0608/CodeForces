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

	var t, n int
	for fmt.Fscan(in, &t); t > 0; t-- {
		fmt.Fscan(in, &n)
		a := make([]int, n)
		for i := range a {
			fmt.Fscan(in, &a[i])
		}
		fmt.Fprintln(out, n)
		for i := 0; i < n; i++ {
			x := a[i]
			if x&(x-1) != 0 {
				x |= x >> 1
				x |= x >> 2
				x |= x >> 4
				x |= x >> 8
				x |= x >> 16
				x = x - (x >> 1)
				x <<= 1
			}
			fmt.Fprintln(out, i+1, x-a[i])
		}
	}
}

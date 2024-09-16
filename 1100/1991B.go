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
next:
	for fmt.Fscan(in, &t); t > 0; t-- {
		fmt.Fscan(in, &n)
		a := make([]int, n-1)
		for i := range a {
			fmt.Fscan(in, &a[i])
		}
		b := make([]int, n)
		for i := range b {
			x, y := 0, 0
			if i > 0 {
				x = a[i-1]
			}
			if i < n-1 {
				y = a[i]
			}
			b[i] = x | y
		}
		for i := 0; i < n-1; i++ {
			if a[i] != b[i]&b[i+1] {
				fmt.Fprintln(out, -1)
				continue next
			}
		}
		for _, x := range b {
			fmt.Fprint(out, x, " ")
		}
		fmt.Fprintln(out)
	}
}

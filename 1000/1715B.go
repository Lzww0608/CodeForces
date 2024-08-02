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

	var t, n, s, k, b int
	for fmt.Fscan(in, &t); t > 0; t-- {
		fmt.Fscan(in, &n, &k, &b, &s)
		a := make([]int, n)
		a[0] = k * b
		s -= k * b
		if s < 0 {
			fmt.Fprintln(out, -1)
			continue
		}
		for i := 0; i < n; i++ {
			cur := min(k-1, s)
			a[i] += cur
			s -= cur
		}
		if s > 0 {
			fmt.Fprintln(out, -1)
			continue
		} else {
			for _, x := range a {
				fmt.Fprint(out, x, " ")
			}
			fmt.Fprintln(out)
		}
	}
}

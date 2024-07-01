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
		mx, idx := -1, -1
		for i, x := range a {
			if x > mx {
				if i > 0 && a[i] > a[i-1] {
					mx, idx = x, i+1
				} else if i < n-1 && a[i] > a[i+1] {
					mx, idx = x, i+1
				}

			}
		}
		fmt.Fprintln(out, idx)

	}
}

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
		b := make([]int, n-1)
		for i := range b {
			fmt.Fscan(in, &b[i])
		}
		a := make([]int, n)
		for i := range a {
			a[i] = -1
		}

		for i := range b {
			if i < n-2 && b[i] < b[i+1] {
				if a[i] == -1 {
					a[i] = b[i]
				}
				a[i+1] = b[i]
			} else {
				if a[i] == -1 {
					a[i] = b[i]
				} else if a[i] < b[i] {
					a[i+1] = b[i]
				}
			}
		}
		a[n-1] = b[n-2]
		for _, x := range a {
			fmt.Fprint(out, x, " ")
		}
		fmt.Fprintln(out)
	}
}

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
		b := make([]int, n)
		c := make([]int, n)
		for i := range a {
			fmt.Fscan(in, &a[i])
		}
		for i := range b {
			fmt.Fscan(in, &b[i])
		}
		for i := range c {
			fmt.Fscan(in, &c[i])
		}
		pre := -1
		for i := 0; i < n; i++ {
			if i == n-1 {
				x := c[n-1]
				if x == pre || x == a[0] {
					x = b[n-1]
					if x == pre || x == a[0] {
						x = a[n-1]
					}
				}
				fmt.Fprint(out, x, " ")
				break
			}

			if pre == a[i] {
				fmt.Fprint(out, b[i], " ")
				pre = b[i]
			} else {
				fmt.Fprint(out, a[i], " ")
				pre = a[i]
			}
		}
		fmt.Fprintln(out)
	}
}

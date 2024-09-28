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
		a := make([]int, n)
		for i := range a {
			fmt.Fscan(in, &a[i])
		}
		if a[0] == 1 {
			fmt.Fprint(out, n+1, " ")
			for i := 1; i <= n; i++ {
				fmt.Fprint(out, i, " ")
			}
			fmt.Fprintln(out)
			continue next
		} else if a[n-1] == 0 {
			for i := 1; i <= n; i++ {
				fmt.Fprint(out, i, " ")
			}
			fmt.Fprintln(out, n+1)
			continue next
		}

		for i := 0; i < n-1; i++ {
			if a[i] == 0 && a[i+1] == 1 {
				for j := 1; j <= i+1; j++ {
					fmt.Fprint(out, j, " ")
				}
				fmt.Fprint(out, n+1, " ")
				for j := i + 2; j <= n; j++ {
					fmt.Fprint(out, j, " ")
				}
				fmt.Fprintln(out)
				continue next
			}
		}
		fmt.Fprintln(out, -1)
	}
}

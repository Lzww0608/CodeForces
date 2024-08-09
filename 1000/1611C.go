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
		if n != a[0] && n != a[n-1] {
			fmt.Fprintln(out, -1)
			continue
		}
		for i := n - 1; i >= 0; i-- {
			if a[i] != n {
				fmt.Fprint(out, a[i], " ")
			}
		}
		fmt.Fprintln(out, n)

	}
}

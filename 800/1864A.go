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

	var t, x, y, n int
	for fmt.Fscan(in, &t); t > 0; t-- {
		fmt.Fscan(in, &x, &y, &n)
		if (n-1)*(n)/2 > y-x {
			fmt.Fprintln(out, -1)
			continue
		}
		a := make([]int, n)
		a[0], a[n-1] = x, y
		for i, j := n-2, 1; i > 0; i, j = i-1, j+1 {
			a[i] = a[i+1] - j
		}
		for _, x := range a {
			fmt.Fprint(out, x, " ")
		}
		fmt.Fprintln(out)

	}
}

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
		if n == 1 {
			fmt.Fprintln(out, 1)
			continue
		} else if n == 2 {
			fmt.Fprintln(out, 1, 2)
			continue
		}

		a := make([]int, n)
		a[0], a[n-1] = 2, 3
		a[n/2] = 1
		x := 4
		for i := 1; i < n-1; i++ {
			if i != n/2 {
				a[i] = x
				x++
			}
		}
		for _, x := range a {
			fmt.Fprint(out, x, " ")
		}
		fmt.Fprintln(out)
	}
}

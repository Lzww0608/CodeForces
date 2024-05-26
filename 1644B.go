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
			a[i] = n - i
		}
		for i := 0; i < n; i++ {
			for _, x := range a {
				fmt.Fprint(out, x, " ")
			}
			fmt.Fprintln(out)
			if i < n-1 {
				a[n-i-2], a[n-i-1] = a[n-i-1], a[n-i-2]
			}

		}

	}
}

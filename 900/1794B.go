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
			a[i] = max(a[i], 2)
		}
		for i := 0; i < n-1; i++ {
			if a[i+1]%a[i] == 0 {
				a[i+1]++
			}

		}
		for _, x := range a {
			fmt.Fprint(out, x, " ")
		}
		fmt.Fprintln(out)

	}
}

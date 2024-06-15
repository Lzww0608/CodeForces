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
		if a[0] == a[n-1] {
			fmt.Fprintln(out, "NO")
		} else {
			fmt.Fprintln(out, "YES")
			fmt.Fprint(out, "R")
			fmt.Fprint(out, "B")
			for i := 2; i < n; i++ {
				fmt.Fprint(out, "R")
			}
			fmt.Fprintln(out)
		}

	}
}

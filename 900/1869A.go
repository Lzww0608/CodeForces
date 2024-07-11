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

	var t, n, x int
	for fmt.Fscan(in, &t); t > 0; t-- {
		fmt.Fscan(in, &n)
		for i := 0; i < n; i++ {
			fmt.Fscan(in, &x)
		}
		if n&1 == 0 {
			fmt.Fprintln(out, 2)
			fmt.Fprintln(out, 1, n)
			fmt.Fprintln(out, 1, n)
		} else {
			fmt.Fprintln(out, 4)
			fmt.Fprintln(out, 1, n)
			fmt.Fprintln(out, 1, n-1)
			fmt.Fprintln(out, 2, n)
			fmt.Fprintln(out, 2, n)
		}
	}
}

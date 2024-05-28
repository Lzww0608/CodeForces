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
		if n&1 == 0 {
			for i := n - 2; i >= 1; i-- {
				fmt.Fprint(out, i, " ")
			}
			fmt.Fprintln(out, n-1, n)
		} else {
			fmt.Fprint(out, 1, " ")
			for i := n - 2; i >= 2; i-- {
				fmt.Fprint(out, i, " ")
			}
			fmt.Fprintln(out, n-1, n)
		}
	}
}

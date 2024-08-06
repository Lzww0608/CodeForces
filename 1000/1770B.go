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

	var t, n, k int
	for fmt.Fscan(in, &t); t > 0; t-- {
		fmt.Fscan(in, &n, &k)
		for l, r := 1, n; l <= r; l, r = l+1, r-1 {
			if l == r {
				fmt.Fprint(out, l)
			} else {
				fmt.Fprint(out, r, l, " ")
			}
		}
		fmt.Fprintln(out)
	}
}

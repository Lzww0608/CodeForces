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
		mid := (n + 1) >> 1
		r := n
		for i := 0; i < n; i++ {
			if i&1 == 0 {
				fmt.Fprint(out, mid, " ")
				mid--
			} else {
				fmt.Fprint(out, r, " ")
				r--
			}

		}
		fmt.Fprintln(out)
	}
}

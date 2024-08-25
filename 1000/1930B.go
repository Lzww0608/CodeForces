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
		l, r := 1, n
		for l <= r {
			if l <= r {
				fmt.Fprint(out, r, " ")
				r--
			}
			if l <= r {
				fmt.Fprint(out, l, " ")
				l++
			}
		}
		fmt.Fprintln(out)
	}
}

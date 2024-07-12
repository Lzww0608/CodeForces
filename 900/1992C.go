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

	var t, n, m, k int
	for fmt.Fscan(in, &t); t > 0; t-- {
		fmt.Fscan(in, &n, &k, &m)
		i := n
		for i >= min(k+1, m) {
			fmt.Fprint(out, i, " ")
			i--
		}

		for j := 1; j <= i; j++ {
			fmt.Fprint(out, j, " ")
		}
		fmt.Fprintln(out)
	}
}

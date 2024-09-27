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
		for k > 3 {
			k--
			n--
			fmt.Fprint(out, 1, " ")
		}
		if n&1 == 1 {
			fmt.Fprintln(out, 1, n/2, n/2)
		} else if n&2 != 0 {
			fmt.Fprintln(out, 2, n/2-1, n/2-1)
		} else {
			fmt.Fprintln(out, n/2, n/4, n/4)
		}
	}
}

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
		if k != 1 && k != n {
			fmt.Fprintln(out, -1)
			continue
		}

		for i := 1; i < n; i++ {
			fmt.Fprint(out, 1, " ")
		}
		if k == 1 {
			fmt.Fprint(out, 2, " ")
		} else {
			fmt.Fprint(out, 1, " ")
		}

		fmt.Fprintln(out)
	}
}

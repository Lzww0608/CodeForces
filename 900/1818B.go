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
		if n == 1 {
			fmt.Fprintln(out, 1)
			continue
		}

		if n&1 == 1 {
			fmt.Fprintln(out, -1)
			continue
		}
		for i := 1; i <= n; i += 2 {
			if i == n {
				fmt.Fprint(out, i, " ")
				break
			}
			fmt.Fprint(out, i+1, " ", i, " ")
		}
		fmt.Fprintln(out)
	}
}

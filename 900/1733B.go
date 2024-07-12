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

	var t, n, x, y int
	for fmt.Fscan(in, &t); t > 0; t-- {
		fmt.Fscan(in, &n, &x, &y)
		if y > x {
			x, y = y, x
		}
		if y > 0 || x == 0 || (n-1)%x != 0 {
			fmt.Fprintln(out, -1)
			continue
		}

		for i := 2; i <= n; i += x {
			for j := 1; j <= x; j++ {
				fmt.Fprint(out, i, " ")
			}
		}
		fmt.Fprintln(out)
	}
}

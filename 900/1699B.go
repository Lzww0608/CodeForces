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

	var t, n, m int
	for fmt.Fscan(in, &t); t > 0; t-- {
		fmt.Fscan(in, &n, &m)
		for i := 1; i <= n; i++ {
			for j := 1; j <= m; j++ {
				if (i%4 <= 1) == (j%4 <= 1) {
					fmt.Fprint(out, 1, " ")
				} else {
					fmt.Fprint(out, 0, " ")
				}
			}
			fmt.Fprintln(out)
		}
	}
}

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
		fmt.Fprint(out, "W")
		for j := 1; j < m; j++ {
			fmt.Fprint(out, "B")
		}
		fmt.Fprintln(out)

		for i := 1; i < n; i++ {
			for j := 0; j < m; j++ {
				fmt.Fprint(out, "B")
			}
			fmt.Fprintln(out)
		}
	}
}

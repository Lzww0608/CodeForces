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
		for i := 0; i < n; i++ {
			for j := 0; j <= i; j++ {
				if j == 0 || j == i {
					fmt.Fprint(out, 1, " ")
				} else {
					fmt.Fprint(out, 0, " ")
				}

			}
			fmt.Fprintln(out)
		}
	}
}

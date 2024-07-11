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
		if n&1 == 0 {
			for i := 1; i <= n; i++ {
				for j := 1; j <= n; j++ {
					if i == j || i == n-j+1 {
						fmt.Fprint(out, 1, " ")
					} else {
						fmt.Fprint(out, 0, " ")
					}

				}
				fmt.Fprintln(out)
			}
		} else {
			for i := 1; i <= n; i++ {
				for j := 1; j <= n; j++ {
					if i == j || i+1 == j || i-1 == j {
						fmt.Fprint(out, 1, " ")
					} else {
						fmt.Fprint(out, 0, " ")
					}

				}
				fmt.Fprintln(out)
			}
		}

	}
}

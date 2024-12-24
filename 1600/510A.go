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

	var n, m int
	fmt.Fscan(in, &n, &m)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if i&1 == 0 {
				fmt.Fprint(out, "#")
			} else {
				if (i/2)&1 == 0 {
					if j == m-1 {
						fmt.Fprint(out, "#")
					} else {
						fmt.Fprint(out, ".")
					}
				} else {
					if j == 0 {
						fmt.Fprint(out, "#")
					} else {
						fmt.Fprint(out, ".")
					}
				}
			}
		}
		fmt.Fprintln(out)
	}
}

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

	var n, m, x, y int
	fmt.Fscan(in, &n, &m, &x, &y)
	fmt.Fprintln(out, x, y)
	for j := 1; j <= m; j++ {
		if j != y {
			fmt.Fprintln(out, x, j)
		}
	}

	k := 0
	for i := 1; i <= n; i++ {
		if i != x {
			if k&1 == 1 {
				for j := 1; j <= m; j++ {
					fmt.Fprintln(out, i, j)
				}
			} else {
				for j := m; j >= 1; j-- {
					fmt.Fprintln(out, i, j)
				}
			}
			k++
		}
	}
}

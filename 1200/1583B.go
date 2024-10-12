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
	var a, b, c int
	for fmt.Fscan(in, &t); t > 0; t-- {
		fmt.Fscan(in, &n, &m)
		vis := make([]bool, n)
		for i := 0; i < m; i++ {
			fmt.Fscan(in, &a, &b, &c)
			vis[b-1] = true
		}
		star := 0
		for i := range vis {
			if !vis[i] {
				star = i + 1
				break
			}
		}
		for i := 1; i <= n; i++ {
			if i != star {
				fmt.Fprintln(out, i, star)
			}
		}
	}
}

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
	cnt := make([]int, m+1)
	for range n {
		var x int
		fmt.Fscan(in, &x)
		cnt[x]++
	}

	f := [3][3]int{}
	for x := 1; x <= m; x++ {
		g := [3][3]int{}
		for i := range 3 {
			for j := range 3 {
				for k := range 3 {
					if i+j+k > cnt[x] {
						continue
					}
					g[j][k] = max(g[j][k], f[i][j]+(cnt[x]-i-j-k)/3+k)
				}
			}
		}

		f = g
	}
	fmt.Fprintln(out, f[0][0])
}

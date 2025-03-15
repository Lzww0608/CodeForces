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

	var t int
	for fmt.Fscan(in, &t); t > 0; t-- {
		solve(in, out)
		fmt.Fprintln(out)
	}
}

func solve(in *bufio.Reader, out *bufio.Writer) {
	var n, m int
	fmt.Fscan(in, &n, &m)
	g := make([]int, n)
	edges := make([][2]int, m)

	for i := 0; i < m; i++ {
		var u, v int
		fmt.Fscan(in, &u, &v)
		u--
		v--
		edges[i] = [2]int{u, v}
		g[u]++
		g[v]++
	}

	ans := make([]int, m)
	for i := 0; i < n; i++ {
		if g[i] == n-1 {
			continue
		}
		for j := 0; j < m; j++ {
			if edges[j][0] == i || edges[j][1] == i {
				ans[j] = 1
			} else {
				ans[j] = 2
			}
		}
		fmt.Fprintln(out, 2)
		for _, v := range ans {
			fmt.Fprint(out, v, " ")
		}
		return
	}

	cnt := 0
	fmt.Fprintln(out, 3)
	for i := 0; i < m; i++ {
		if edges[i][0] == 0 || edges[i][1] == 0 {
			if cnt == 0 {
				ans[i] = 1
			} else {
				ans[i] = 2
			}
			cnt++
		} else {
			ans[i] = 3
		}
	}

	for _, v := range ans {
		fmt.Fprint(out, v, " ")
	}
	return

}

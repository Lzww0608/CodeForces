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
	var s string
	fmt.Fscan(in, &s)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = int(s[i] - 'a')
	}
	p := [][]int{{0, 1, 2}, {0, 2, 1}, {1, 2, 0}, {1, 0, 2}, {2, 0, 1}, {2, 1, 0}}
	sum := make([][]int, 6)
	for i := range sum {
		sum[i] = make([]int, n+1)
	}
	for j, v := range p {
		for i := 0; i < n; i++ {
			if a[i] == v[i%3] {
				sum[j][i+1] = sum[j][i]
			} else {
				sum[j][i+1] = sum[j][i] + 1
			}
		}
	}

	var l, r int
	for i := 0; i < m; i++ {
		fmt.Fscan(in, &l, &r)
		l--
		ans := n
		for _, v := range sum {
			ans = min(ans, v[r]-v[l])
		}
		fmt.Fprintln(out, ans)
	}
}

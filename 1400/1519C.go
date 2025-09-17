package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

const N int = 10

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var t int
	for fmt.Fscan(in, &t); t > 0; t-- {
		solve(in, out)
	}
}

func solve(in *bufio.Reader, out *bufio.Writer) {
	var n int
	fmt.Fscan(in, &n)
	u := make([]int, n)
	s := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Fscan(in, &u[i])
		u[i]--
	}
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &s[i])
	}

	m := make([][]int, n)
	for i := 0; i < n; i++ {
		x, y := u[i], s[i]
		m[x] = append(m[x], y)
	}

	for k := 0; k < n; k++ {
		sort.Slice(m[k], func(i, j int) bool {
			return m[k][i] > m[k][j]
		})
	}

	pre := make([][]int, n)
	for i := 0; i < n; i++ {
		pre[i] = make([]int, len(m[i])+1)
		for j := range m[i] {
			pre[i][j+1] = pre[i][j] + m[i][j]
		}
	}

	ans := make([]int, n)
	for i := 0; i < n; i++ {
		for k := 1; k <= len(m[i]); k++ {
			ans[k-1] += pre[i][len(m[i])/k*k]
		}
	}

	for _, x := range ans {
		fmt.Fprint(out, x, " ")
	}
	fmt.Fprintln(out)
}

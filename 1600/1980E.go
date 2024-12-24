package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

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
	var n, m int
	fmt.Fscan(in, &n, &m)
	a := make([][]int, n)
	b := make([][]int, n)
	c := make([][]int, n)
	id1 := make([]int, m)
	id2 := make([]int, m)
	for i := 0; i < n; i++ {
		a[i] = make([]int, m)
		c[i] = make([]int, m)
		for j := 0; j < m; j++ {
			id1[j] = j
			id2[j] = j
			fmt.Fscan(in, &a[i][j])
		}
	}
	pos := 0
	for i := 0; i < n; i++ {
		b[i] = make([]int, m)
		for j := 0; j < m; j++ {
			fmt.Fscan(in, &b[i][j])
			if b[i][j] == a[0][0] {
				pos = i
			}
		}
	}

	sort.Slice(id1, func(i, j int) bool {
		return a[0][id1[i]] < a[0][id1[j]]
	})
	sort.Slice(id2, func(i, j int) bool {
		return b[pos][id2[i]] < b[pos][id2[j]]
	})

	for j := 0; j < m; j++ {
		for i := 0; i < n; i++ {
			c[i][j] = a[i][id1[j]]
		}
	}
	for j := 0; j < m; j++ {
		for i := 0; i < n; i++ {
			a[i][j] = b[i][id2[j]]
		}
	}

	sort.Slice(a, func(i, j int) bool {
		return a[i][0] < a[j][0]
	})

	sort.Slice(c, func(i, j int) bool {
		return c[i][0] < c[j][0]
	})

	//fmt.Fprintln(out, a, c)

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if a[i][j] != c[i][j] {
				fmt.Fprintln(out, "NO")
				return
			}
		}
	}
	fmt.Fprintln(out, "YES")
}

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
	}
}

func solve(in *bufio.Reader, out *bufio.Writer) {
	var n, m int
	fmt.Fscan(in, &n, &m)
	a := make([]int, n)
	b := make([]int, n)
	c := make([]int, m)
	wants := make([][]int, n+1)
	needs := make([][]int, n+1)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])

	}
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &b[i])
		wants[b[i]] = append(wants[b[i]], i)
		if a[i] != b[i] {
			needs[b[i]] = append(needs[b[i]], i)
		}
	}
	for i := 0; i < m; i++ {
		fmt.Fscan(in, &c[i])
	}

	if len(wants[c[m-1]]) == 0 {
		fmt.Fprintln(out, "NO")
		return
	}

	ans := make([]int, m)
	for i := m - 1; i >= 0; i-- {
		x := c[i]
		if len(needs[x]) != 0 {
			ans[i] = needs[x][0] + 1
			needs[x] = needs[x][1:]
		} else if len(wants[x]) != 0 {
			ans[i] = wants[x][0] + 1
		} else {
			ans[i] = ans[m-1]
		}
	}

	for _, v := range needs {
		if len(v) != 0 {
			fmt.Fprintln(out, "NO")
			return
		}
	}

	for _, x := range ans {
		if x == 0 {
			fmt.Fprintln(out, "NO")
			return
		}
	}
	fmt.Fprintln(out, "YES")
	for _, x := range ans {
		fmt.Fprint(out, x, " ")
	}
	fmt.Fprintln(out)
}

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
	a := make([]int, n)
	c := make([]int, m)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
		a[i]--
	}
	for i := 0; i < m; i++ {
		fmt.Fscan(in, &c[i])
	}

	ans := 0
	sort.Ints(a)
	j := 0
	for k, i := range a {
		if n-k-1 >= i || j == m {
			ans += c[i]
		} else {
			ans += c[j]
			j++
		}
	}

	fmt.Fprintln(out, ans)
}

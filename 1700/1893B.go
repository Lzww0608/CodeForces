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
	b := make([]int, m)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
	}
	for i := 0; i < m; i++ {
		fmt.Fscan(in, &b[i])
	}
	sort.Slice(b, func(i, j int) bool {
		return b[i] > b[j]
	})
	j := 0
	for i := 0; i < n; i++ {
		for j < m && b[j] >= a[i] {
			fmt.Fprint(out, b[j], " ")
			j++
		}
		fmt.Fprint(out, a[i], " ")
	}
	for j < m {
		fmt.Fprint(out, b[j], " ")
		j++
	}
	fmt.Fprintln(out)
}

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
	var n int
	fmt.Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		fmt.Fscan(in, &a[i])
	}
	deg := make([]int, n)
	for i := 0; i < n-1; i++ {
		var u, v int
		fmt.Fscan(in, &u, &v)
		u--
		v--
		deg[u]++
		deg[v]++
	}

	ans := 0
	b := []int{}
	for i := 0; i < n; i++ {
		for j := 1; j < deg[i]; j++ {
			b = append(b, a[i])
		}
		ans += a[i]
	}

	sort.Slice(b, func(i, j int) bool {
		return b[i] > b[j]
	})
	for _, x := range b {
		fmt.Fprint(out, ans, " ")
		ans += x
	}
	fmt.Fprintln(out, ans)
	return
}

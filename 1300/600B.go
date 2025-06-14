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

	var n, m int
	fmt.Fscan(in, &n, &m)
	a := make([]int, n)
	b := make([]int, m)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
	}

	id := make([]int, m)
	for i := 0; i < m; i++ {
		fmt.Fscan(in, &b[i])
		id[i] = i
	}
	sort.Ints(a)
	sort.Slice(id, func(i, j int) bool {
		return b[id[i]] < b[id[j]]
	})

	ans := make([]int, m)
	i := 0
	for _, j := range id {
		for i < n && a[i] <= b[j] {
			i++
		}
		ans[j] = i
	}
	for _, x := range ans {
		fmt.Fprintf(out, "%d ", x)
	}
}

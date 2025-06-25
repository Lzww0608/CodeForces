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

	var n, d int
	fmt.Fscan(in, &n, &d)
	ans := 0
	a := make([]int, n)
	b := make([]int, n)
	id := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i], &b[i])
		id[i] = i
	}

	sort.Slice(id, func(i, j int) bool {
		return a[id[i]] < a[id[j]]
	})

	cur := 0
	for l, r := 0, 0; r < n; r++ {
		cur += b[id[r]]
		for a[id[r]]-a[id[l]] >= d {
			cur -= b[id[l]]
			l++
		}
		ans = max(ans, cur)
	}

	fmt.Fprintln(out, ans)
}

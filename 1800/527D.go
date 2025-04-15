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

	var n int
	fmt.Fscan(in, &n)
	a := make([]int, n)
	b := make([]int, n)
	id := make([]int, n)
	for i := 0; i < n; i++ {
		id[i] = i
		fmt.Fscan(in, &a[i], &b[i])
	}

	sort.Slice(id, func(i, j int) bool {
		return a[id[i]]+b[id[i]] < a[id[j]]+b[id[j]]
	})

	ans := 1
	pre := 0
	for i := 1; i < n; i++ {
		if a[id[i]]-b[id[i]] >= a[id[pre]]+b[id[pre]] {
			ans++
			pre = i
		}
	}
	fmt.Fprintln(out, ans)
}

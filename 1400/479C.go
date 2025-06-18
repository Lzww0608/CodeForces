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
		return a[id[i]] < a[id[j]] || a[id[i]] == a[id[j]] && b[id[i]] < b[id[j]]
	})

	ans := -1
	for _, i := range id {
		if ans <= b[i] {
			ans = b[i]
		} else {
			ans = a[i]
		}
	}

	fmt.Fprintln(out, ans)
}

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

	var n, k int
	fmt.Fscan(in, &n, &k)

	a := make([]int, n)
	id := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
		id[i] = i
	}

	sort.Slice(id, func(i, j int) bool {
		return a[id[i]] < a[id[j]]
	})

	cur := 0
	ans := []int{}
	for _, i := range id {
		if cur += a[i]; cur <= k {
			ans = append(ans, i+1)
		} else {
			break
		}
	}

	fmt.Fprintln(out, len(ans))
	for _, i := range ans {
		fmt.Fprint(out, i, " ")
	}
}

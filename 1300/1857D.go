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
	b := make([]int, n)
	id := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
	}
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &b[i])
		id[i] = i
	}
	sort.Slice(id, func(i, j int) bool {
		return a[id[i]]-b[id[i]] > a[id[j]]-b[id[j]]
	})

	ans := []int{id[0]}
	for i := 1; i < n; i++ {
		if a[id[i]]-b[id[i]] != a[id[i-1]]-b[id[i-1]] {
			break
		}
		ans = append(ans, id[i])
	}
	sort.Ints(ans)
	fmt.Fprintln(out, len(ans))
	for _, x := range ans {
		fmt.Fprint(out, x+1, " ")
	}
	fmt.Fprintln(out)
}

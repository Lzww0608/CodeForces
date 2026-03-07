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
	id := make([]int, n)
	ans := make([]int, n)
	sum := 0
	for i := range a {
		id[i] = i
		fmt.Fscan(in, &a[i])
		sum += a[i]
	}

	sort.Slice(id, func(i, j int) bool {
		return a[id[i]] < a[id[j]]
	})

	cur := 0
	for j, i := range id {
		x := a[i]
		ans[i] = sum - (n-j)*x + j*x - cur
		cur += x
		sum -= x
	}

	for _, x := range ans {
		fmt.Fprintf(out, "%d ", x+n)
	}
	fmt.Fprintln(out)
}

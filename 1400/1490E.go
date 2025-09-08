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

	type pair struct {
		x, i int
	}
	a := make([]pair, n)
	for i := range n {
		fmt.Fscan(in, &a[i].x)
		a[i].i = i + 1
	}

	sort.Slice(a, func(i, j int) bool { return a[i].x > a[j].x })

	sum := 0
	for _, v := range a {
		sum += v.x
	}

	ans := []int{}
	pre := 0
	for _, v := range a {
		if sum >= pre {
			ans = append(ans, v.i)
		} else {
			break
		}
		sum -= v.x
		pre = v.x
	}
	sort.Ints(ans)
	fmt.Fprintln(out, len(ans))
	for _, x := range ans {
		fmt.Fprintf(out, "%d ", x)
	}
	fmt.Fprintln(out, "")
}

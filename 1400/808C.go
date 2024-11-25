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

	var n, w int
	fmt.Fscan(in, &n, &w)
	a := make([]int, n)
	p := make([]int, n)
	sum := 0
	for i := 0; i < n; i++ {
		p[i] = i
		fmt.Fscan(in, &a[i])
		sum += (a[i] + 1) / 2
	}

	if sum > w {
		fmt.Fprintln(out, -1)
		return
	}
	t := w % sum
	q := w / sum
	ans := make([]int, n)
	sort.Slice(p, func(i, j int) bool {
		return a[p[i]] > a[p[j]]
	})

	for _, i := range p {
		ans[i] = q * ((a[i] + 1) / 2)
		if t > 0 {
			tmp := min(t, a[i]-ans[i])
			ans[i] += tmp
			t -= tmp
		}
	}
	for _, x := range ans {
		fmt.Fprint(out, x, " ")
	}
}

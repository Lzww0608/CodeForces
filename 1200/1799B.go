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
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
		b[i] = a[i]
	}
	sort.Ints(b)
	if b[n-1] == b[0] {
		fmt.Fprintln(out, 0)
		return
	} else if b[0] == 1 {
		fmt.Fprintln(out, -1)
		return
	}
	ans := [][]int{}

	find := func() (mn, mx int) {
		for i := range a {
			if a[i] > a[mx] {
				mx = i
			}
			if a[i] < a[mn] {
				mn = i
			}
		}
		return
	}

	for {
		mn, mx := find()
		if mn == mx {
			break
		}
		ans = append(ans, []int{mx + 1, mn + 1})
		a[mx] = (a[mx] + a[mn] - 1) / a[mn]
	}
	fmt.Fprintln(out, len(ans))
	for _, v := range ans {
		fmt.Fprintln(out, v[0], v[1])
	}
}

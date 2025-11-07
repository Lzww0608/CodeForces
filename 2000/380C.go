package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type pair struct {
	l, r, i int
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var s string
	fmt.Fscan(in, &s)
	st := []int{}

	n := len(s)
	f := make([]int, n+1)

	update := func(i, v int) {
		for ; i <= n; i += i & (-i) {
			f[i] += v
		}
	}

	query := func(i int) int {
		res := 0
		for ; i > 0; i -= i & (-i) {
			res += f[i]
		}

		return res
	}

	var m int
	fmt.Fscan(in, &m)
	a := make([]pair, m)
	for i := range m {
		fmt.Fscan(in, &a[i].l, &a[i].r)
		a[i].l--
		a[i].i = i
	}

	sort.Slice(a, func(i, j int) bool {
		return a[i].r < a[j].r
	})

	j := 0
	ans := make([]int, m)
	for i := range n {
		if s[i] == '(' {
			st = append(st, i)
		} else {
			if len(st) > 0 {
				k := st[len(st)-1]
				st = st[:len(st)-1]
				update(k+1, 2)
			}
		}

		for j < m && a[j].r == i+1 {
			ans[a[j].i] = query(a[j].r) - query(a[j].l)
			j++
		}
	}

	for _, x := range ans {
		fmt.Fprintln(out, x)
	}
}

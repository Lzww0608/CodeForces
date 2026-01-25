package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"sort"
)

// https://codeforces.com/problemset/problem/86/D
func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, t int
	fmt.Fscan(in, &n, &t)
	a := make([]int, n)
	sz := int(math.Sqrt(float64(n)))
	pos := make([]int, n)
	for i := range a {
		pos[i] = i / sz
		fmt.Fscan(in, &a[i])
	}

	q := make([]Query, t)
	for i := range t {
		var l, r int
		fmt.Fscan(in, &l, &r)
		q[i] = Query{l - 1, r - 1, i}
	}

	sort.Slice(q, func(i, j int) bool {
		if pos[q[i].l] != pos[q[j].l] {
			return pos[q[i].l] < pos[q[j].l]
		}

		if pos[q[i].l]&1 == 0 {
			return q[i].r < q[j].r
		} else {
			return q[i].r > q[j].r
		}
	})

	ans := make([]int, t)
	l, r := 0, -1
	cur := 0
	cnt := make([]int, slices.Max(a)+1)

	add := func(v int) {
		cur -= cnt[v] * cnt[v] * v
		cnt[v]++
		cur += cnt[v] * cnt[v] * v
	}

	remove := func(v int) {
		cur -= cnt[v] * cnt[v] * v
		cnt[v]--
		cur += cnt[v] * cnt[v] * v
	}

	for _, v := range q {
		for r > v.r {
			remove(a[r])
			r--
		}
		for r < v.r {
			r++
			add(a[r])
		}
		for l < v.l {
			remove(a[l])
			l++
		}
		for l > v.l {
			l--
			add(a[l])
		}
		ans[v.id] = cur
	}

	for _, x := range ans {
		fmt.Fprintln(out, x)
	}
}

type Query struct {
	l, r, id int
}

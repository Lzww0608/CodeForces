package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

const inf int = math.MaxInt32

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, q int
	fmt.Fscan(in, &n, &q)
	a := make([]int, n+1)
	cnt, mx := 0, 0
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i])
		mx = max(mx, a[i])
		if a[i] == 0 {
			a[i] = inf
			cnt++
		}
	}

	if mx < q && cnt == 0 {
		fmt.Fprintln(out, "NO")
		return
	}

	minv := make([]int, n*4+10)
	vis := make(map[int]bool)
	m := make(map[int]int)
	var ql, qr int

	var build func(int, int, int)
	build = func(o, l, r int) {
		if l == r {
			minv[o] = a[l]
			return
		}
		mid := l + ((r - l) >> 1)
		build(o*2, l, mid)
		build(o*2+1, mid+1, r)
		minv[o] = min(minv[o*2], minv[o*2+1])
	}

	var query func(int, int, int) int
	query = func(o, l, r int) int {
		res := inf
		mid := l + ((r - l) >> 1)
		if ql <= l && r <= qr {
			return minv[o]
		}
		if ql <= mid {
			res = min(res, query(o*2, l, mid))
		}
		if mid < qr {
			res = min(res, query(o*2+1, mid+1, r))
		}
		return res
	}

	build(1, 1, n)
	for i := 1; i <= n; i++ {
		if a[i] == inf {
			continue
		}
		if vis[a[i]] {
			ql = m[a[i]]
			qr = i
			if query(1, 1, n) < a[i] {
				fmt.Fprintln(out, "NO")
				return
			}
		}
		vis[a[i]] = true
		m[a[i]] = i
	}

	for i := 1; i <= n; i++ {
		if a[i] != inf {
			continue
		}

		if mx < q {
			a[i] = q
		} else {
			a[i] = a[i-1]
		}
		mx = max(mx, a[i])
	}

	for i := n - 1; i >= 1; i-- {
		if a[i] != inf && a[i] != 0 {
			continue
		}

		if mx < q {
			a[i] = q
		} else {
			a[i] = a[i+1]
		}
		mx = max(mx, a[i])
	}

	fmt.Fprintln(out, "YES")
	for i := 1; i <= n; i++ {
		fmt.Fprintf(out, "%d ", a[i])
	}
}

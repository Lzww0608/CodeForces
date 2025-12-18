package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, a int
	fmt.Fscan(in, &n, &a)
	d := make([]int, n)
	c := make([]int, n)
	b := make([]int, n)
	pre := make([]int, n-1)

	ans := 0
	left := make([]int, n)
	right := make([]int, n)
	st1 := []int{}
	st2 := []int{}
	for i := range n {
		fmt.Fscan(in, &d[i], &c[i])
		b[i] = a - c[i]
		ans = max(ans, b[i])
		if i > 0 {
			pre[i-1] = (d[i-1] - d[i]) * (d[i-1] - d[i])
		}
	}
	if n == 1 {
		fmt.Fprintln(out, ans)
		return
	}

	for i := range n - 1 {
		for len(st1) > 0 && pre[st1[len(st1)-1]] <= pre[i] {
			st1 = st1[:len(st1)-1]
		}
		if len(st1) > 0 {
			left[i] = st1[len(st1)-1]
		} else {
			left[i] = -1
		}
		st1 = append(st1, i)
	}

	for i := n - 2; i >= 0; i-- {
		for len(st2) > 0 && pre[st2[len(st2)-1]] <= pre[i] {
			st2 = st2[:len(st2)-1]
		}
		if len(st2) > 0 {
			right[i] = st2[len(st2)-1]
		} else {
			right[i] = n - 1
		}
		st2 = append(st2, i)
	}

	sum := make([]int, n+1)
	for i := range n {
		sum[i+1] = sum[i] + b[i]
	}
	w := bits.Len(uint(n + 1))
	st := make([][2][]int, w)
	for i := range st {
		st[i][0] = make([]int, n+1)
		st[i][1] = make([]int, n+1)
	}
	for i := range n + 1 {
		st[0][0][i] = sum[i]
		st[0][1][i] = sum[i]
	}

	for i := 1; i < w; i++ {
		for j := range n - (1 << i) + 2 {
			st[i][0][j] = min(st[i-1][0][j], st[i-1][0][j+(1<<(i-1))])
			st[i][1][j] = max(st[i-1][1][j], st[i-1][1][j+(1<<(i-1))])
		}
	}

	query := func(l, r int) (int, int) {
		if l > r {
			return 0, 0
		}
		k := bits.Len(uint(r-l)) - 1
		mn := min(st[k][0][l], st[k][0][r-(1<<k)])
		mx := max(st[k][1][l], st[k][1][r-(1<<k)])
		return mx, mn
	}

	for i := range n - 1 {
		l := left[i]
		r := right[i]
		mx, _ := query(i+2, r+2)
		_, mn := query(l+1, i+1)
		ans = max(ans, mx-mn-pre[i])
	}

	fmt.Fprintln(out, ans)
}

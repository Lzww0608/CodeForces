package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
	}

	l := make([]int, n)
	r := make([]int, n)
	for i := 0; i < n; i++ {
		l[i] = -1
		r[i] = n
	}
	st := []int{-1}
	for i, x := range a {
		for len(st) > 1 && a[st[len(st)-1]] >= x {
			cur := st[len(st)-1]
			st = st[:len(st)-1]
			l[cur] = st[len(st)-1]
			r[cur] = i
		}
		st = append(st, i)
	}

	st = []int{n}
	for i := n - 1; i >= 0; i-- {
		x := a[i]
		for len(st) > 1 && a[st[len(st)-1]] >= x {
			cur := st[len(st)-1]
			st = st[:len(st)-1]
			r[cur] = st[len(st)-1]
			l[cur] = i
		}
		st = append(st, i)
	}

	ans := make([]int, n+1)
	for i := 0; i < n; i++ {
		d := r[i] - l[i] - 1
		ans[d] = max(ans[d], a[i])
	}
	ans[n] = slices.Min(a)

	for i := n - 1; i > 0; i-- {
		ans[i] = max(ans[i], ans[i+1])
	}

	for _, x := range ans[1:] {
		fmt.Fprint(out, x, " ")
	}
}

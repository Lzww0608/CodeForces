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

func merge(a []int, l, r int) int {
	if l >= r {
		return 0
	}
	mid := (l + r) / 2
	cnt := merge(a, l, mid) + merge(a, mid+1, r)
	for i, j := l, mid+1; i <= mid; i++ {
		for j <= r && a[j] < a[i] {
			j++
		}
		cnt += j - mid - 1
	}

	tmp := make([]int, r-l+1)
	i, j, k := l, mid+1, 0
	for i <= mid || j <= r {
		if j > r || i <= mid && a[i] < a[j] {
			tmp[k] = a[i]
			i++
		} else {
			tmp[k] = a[j]
			j++
		}
		k++
	}
	copy(a[l:r+1], tmp)
	return cnt
}

func solve(in *bufio.Reader, out *bufio.Writer) {
	var n int
	fmt.Fscan(in, &n)
	a := make([]int, n)
	b := make([]int, n)
	id := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i], &b[i])
		id[i] = i
	}
	sort.Slice(id, func(i, j int) bool {
		return b[id[i]] < b[id[j]]
	})

	c := make([]int, n)
	for i := range n {
		c[i] = a[id[i]]
	}

	fmt.Fprintln(out, merge(c, 0, n-1))
}

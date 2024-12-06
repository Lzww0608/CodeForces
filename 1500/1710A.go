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
	var n, m, k int
	fmt.Fscan(in, &n, &m, &k)
	a := make([]int, k)
	sum := 0
	for i := range a {
		fmt.Fscan(in, &a[i])
		sum += a[i]
	}

	if sum < n*m {
		fmt.Fprintln(out, "NO")
		return
	}
	if k == 1 {
		fmt.Fprintln(out, "YES")
		return
	}
	sort.Slice(a, func(i, j int) bool {
		return a[i] > a[j]
	})

	cnt := 0
	f := false
	for _, x := range a {
		t := x / m
		if t > 1 {
			if n&1 == 0 || t&1 == 1 {
				f = true
			}
			cnt += t
		}
	}
	if cnt >= n && f || cnt > n && a[0]/m > 2 {
		fmt.Fprintln(out, "YES")
		return
	}

	cnt = 0
	f = false
	for _, x := range a {
		t := x / n
		if t > 1 {
			if m&1 == 0 || t&1 == 1 {
				f = true
			}
			cnt += t
		}
	}
	if cnt >= m && f || cnt > m && a[0]/n > 2 {
		fmt.Fprintln(out, "YES")
		return
	}
	fmt.Fprintln(out, "NO")
}

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
	c := make([]int, n)
	for i := range a {
		fmt.Fscan(in, &a[i])
	}
	for i := range b {
		fmt.Fscan(in, &b[i])
	}
	for i := range c {
		fmt.Fscan(in, &c[i])
	}
	sort.Ints(a)
	sort.Ints(b)
	sort.Ints(c)
	ans := 0
	st := make([]int, 0, n)
	for i, j := 0, 0; i < n; i++ {
		for j < n && a[j] < b[i] {
			st = append(st, a[j])
			j++
		}
		b[i] -= st[len(st)-1]
		st = st[:len(st)-1]
	}
	sort.Ints(b)

	for i, x := range c {
		ans += x * b[n-i-1]
	}
	fmt.Fprintln(out, ans)
}

package main

import (
	"bufio"
	"fmt"
	"os"
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
	var n, m, x int
	fmt.Fscan(in, &n)
	a := make([]int, n)
	b := map[int]int{}
	cnt := map[int]int{}
	for i := range a {
		fmt.Fscan(in, &a[i])
	}
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &x)
		b[x]++
		if a[i] == x {
			cnt[a[i]] += 0
		} else {
			cnt[x] += 1
		}
	}
	fmt.Fscan(in, &m)
	d := make([]int, m)
	for i := range d {
		fmt.Fscan(in, &d[i])
	}

	if _, ok := b[d[m-1]]; !ok {
		fmt.Fprintln(out, "NO")
		return
	}
	for i := m - 1; i >= 0; i-- {
		x := d[i]
		if _, ok := b[x]; ok {
			cnt[x]--
		}
	}
	for _, x := range cnt {
		if x > 0 {
			fmt.Fprintln(out, "NO")
			return
		}
	}
	fmt.Fprintln(out, "YES")
}

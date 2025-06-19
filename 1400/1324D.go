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

	var n int
	fmt.Fscan(in, &n)
	a := make([]int, n)
	b := make([]int, n)
	c := make([]int, n)
	d := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
	}
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &b[i])
	}

	ans := 0
	for i := 0; i < n; i++ {
		c[i] = a[i] - b[i]
		d[i] = b[i] - a[i]
		if c[i] > d[i] {
			ans--
		}
	}
	sort.Ints(c)
	sort.Ints(d)

	// a[i] + a[j] > b[i] + b[j]
	// a[i] - b[i] > b[j] - a[j]
	i, j := 0, 0
	for i < n {
		for j < n && d[j] < c[i] {
			j++
		}
		ans += j
		i++
	}

	fmt.Fprintln(out, ans/2)
}

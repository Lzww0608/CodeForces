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

	var n, m int
	fmt.Fscan(in, &n, &m)
	a := make([]int, n)
	b := make([]int, m)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
	}
	for i := 0; i < m; i++ {
		fmt.Fscan(in, &b[i])
	}

	ans := 0
	c := make([]int, n+1)
	k := 1
	for _, x := range b {
		if c[x] == 0 {
			c[x] = k
			k++
		}
	}
	for i := 0; i < m; i++ {
		if i == 0 || b[i] == b[i-1] {
			continue
		}
		cur := 0
		idx := c[b[i]]
		for i := 1; i <= n; i++ {
			if c[i] < idx && c[i] > 0 {
				cur += a[i-1]
			}
		}
		ans += cur
		for i := 1; i <= n; i++ {
			if c[i] > 0 && c[i] < idx {
				c[i]++
			}
		}
		c[b[i]] = 1
	}

	fmt.Fprintln(out, ans)
}

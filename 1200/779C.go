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

	var n, k int
	fmt.Fscan(in, &n, &k)
	a := make([]int, n)
	b := make([]int, n)
	for i := range a {
		fmt.Fscan(in, &a[i])
	}
	for i := range b {
		fmt.Fscan(in, &b[i])
	}
	f := make([]bool, n)
	cnt, ans := 0, 0
	for i := range f {
		if a[i] <= b[i] {
			cnt++
			ans += a[i]
			f[i] = true
		}
	}

	if cnt >= k {
		for i := range f {
			if !f[i] {
				ans += b[i]
			}
		}
		fmt.Fprintln(out, ans)
	} else {
		cnt = k - cnt
		c := make([][]int, 0, n)
		for i := range f {
			if !f[i] {
				c = append(c, []int{i, a[i] - b[i]})
			}
		}
		sort.Slice(c, func(i, j int) bool {
			return c[i][1] < c[j][1]
		})
		for i := 0; i < cnt; i++ {
			ans += a[c[i][0]]
			f[c[i][0]] = true
		}
		for i := range f {
			if !f[i] {
				ans += b[i]
			}
		}
		fmt.Fprintln(out, ans)
	}

}

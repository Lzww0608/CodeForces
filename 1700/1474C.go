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
	n *= 2
	a := make([]int, n)
	for i := range a {
		fmt.Fscan(in, &a[i])
	}
	sort.Ints(a)
next:
	for i := 0; i < n-1; i++ {
		x := a[i] + a[n-1]
		ans := make([][]int, 0, n/2)
		ans = append(ans, []int{a[i], a[n-1]})
		cnt := make(map[int]int)
		for _, y := range a {
			cnt[y]++
		}
		cnt[a[i]]--
		cnt[a[n-1]]--
		p, q, k := a[i], a[n-1], n-2
		for j := 1; j < n/2; j++ {
			x = max(p, q)
			for k >= 0 && cnt[a[k]] == 0 {
				k--
			}
			if k < 0 || a[k] >= x {
				continue next
			} else {
				cnt[a[k]]--
				y := x - a[k]
				if cnt[y] == 0 {
					continue next
				}
				cnt[y]--
				p, q = a[k], x-a[k]
				ans = append(ans, []int{p, q})
			}
		}
		fmt.Fprintln(out, "YES")
		fmt.Fprintln(out, a[i]+a[n-1])
		for _, v := range ans {
			fmt.Fprintln(out, v[0], v[1])
		}
		return
	}
	fmt.Fprintln(out, "No")
}

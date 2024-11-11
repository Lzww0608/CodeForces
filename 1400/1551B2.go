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
	var n, k int
	fmt.Fscan(in, &n, &k)
	a := make([]int, n)
	cnt := make([]int, n)
	ans := make([]int, n)
	var p []int
	for i := range a {
		fmt.Fscan(in, &a[i])
		a[i]--
		cnt[a[i]]++
		if cnt[a[i]] <= k {
			p = append(p, i)
		}
	}

	for len(p)%k != 0 {
		p = p[:len(p)-1]
	}

	sort.Slice(p, func(i, j int) bool {
		return a[p[i]] < a[p[j]]
	})
	//fmt.Fprintln(out, p)
	for i := range p {
		ans[p[i]] = i%k + 1
	}

	for _, x := range ans {
		fmt.Fprint(out, x, " ")
	}
	fmt.Fprintln(out)
}

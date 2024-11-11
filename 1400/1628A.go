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
	var n int
	fmt.Fscan(in, &n)
	a := make([]int, n)
	cnt := map[int]int{}
	p := 0
	for i := range a {
		fmt.Fscan(in, &a[i])
		cnt[a[i]]++
	}
	for cnt[p] > 0 {
		p++
	}
	ans := []int{}

	i := 0
	for i < n && len(cnt) > 0 {
		if p == 0 {
			for i < n {
				ans = append(ans, 0)
				i++
			}
			break
		}
		vis := map[int]bool{}
		for i < n && len(vis) < p {
			x := a[i]
			cnt[x]--
			if cnt[x] == 0 {
				delete(cnt, x)
			}
			if x < p {
				vis[x] = true
			}
			i++
		}
		ans = append(ans, p)
		p = 0
		for cnt[p] > 0 {
			p++
		}
	}
	fmt.Fprintln(out, len(ans))
	for _, x := range ans {
		fmt.Fprint(out, x, " ")
	}
	fmt.Fprintln(out)
}

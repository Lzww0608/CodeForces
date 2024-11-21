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

	var s string
	fmt.Fscanln(in, &s)
	var m int
	fmt.Fscanln(in, &m)
	a := make([]int, m)
	for i := 0; i < m; i++ {
		fmt.Fscan(in, &a[i])
	}
	n := len(s)
	cnt := make([]int, n+1)
	for _, x := range a {
		y := n - x + 1
		x--
		y--
		cnt[x]++
		cnt[y+1]--
	}

	ans := make([]byte, n)
	cur := 0
	for i := 0; i < n; i++ {
		cur += cnt[i]
		if cur&1 == 0 {
			ans[i] = s[i]
		} else {
			ans[n-i-1] = s[i]
		}
	}
	fmt.Fprintln(out, string(ans))
}

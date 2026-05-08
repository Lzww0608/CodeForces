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
	var n, k int
	fmt.Fscan(in, &n, &k)
	a := make([]int, n)
	for i := range a {
		fmt.Fscan(in, &a[i])
	}

	cnt := make([]int, k*2+1)
	pre := make([]int, k*2+2)
	for i := 0; i < n/2; i++ {
		x, y := a[i], a[n-1-i]
		cnt[x+y]++
		pre[min(x, y)+1]++
		pre[max(x, y)+k+1]--
	}

	for i := 1; i < len(pre); i++ {
		pre[i] += pre[i-1]
	}

	ans := n
	for i := 2; i <= 2*k; i++ {
		cur := (pre[i] - cnt[i]) + (n/2-pre[i])*2
		ans = min(ans, cur)
	}
	fmt.Fprintln(out, ans)
}

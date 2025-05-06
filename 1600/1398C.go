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
	var s string
	fmt.Fscan(in, &s)
	pre := make([]int, n+1)
	for i := 0; i < n; i++ {
		pre[i+1] = pre[i] + int(s[i]-'0')
	}

	ans := 0
	cnt := make(map[int]int)
	for i := 0; i <= n; i++ {
		ans += cnt[pre[i]-i]
		cnt[pre[i]-i]++
	}

	fmt.Fprintln(out, ans)
}

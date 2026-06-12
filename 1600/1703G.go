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
	pre := make([]int, n+1)
	for i := range a {
		fmt.Fscan(in, &a[i])
		pre[i+1] = pre[i] + a[i]
	}

	ans := 0
	for i := range n + 1 {
		cur := pre[i] - i*k
		for j := i; j < n && j-i <= 30; j++ {
			t := a[j]
			for range j - i + 1 {
				t /= 2
			}
			cur += t
		}

		ans = max(ans, cur)
		//fmt.Fprintln(out, ans)
	}

	fmt.Fprintln(out, ans)
}

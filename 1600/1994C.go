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
	pre := make([]int, n+1)
	for i := range a {
		fmt.Fscan(in, &a[i])
		pre[i+1] = pre[i] + a[i]
	}

	f := make([]int, n+2)
	ans := 0

	for i := n - 1; i >= 0; i-- {
		j := sort.SearchInts(pre, pre[i]+k+1) - 1
		f[i] = f[j+1] + j - i
		ans += f[i]
	}

	fmt.Fprintln(out, ans)
}

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
	a := make([]int, n)
	for i := range a {
		fmt.Fscan(in, &a[i])
	}
	sort.Ints(a)
	ans := 0
	for i := 0; i < n; i++ {
		if i-2 >= 0 {
			cur := a[i] - a[i-1] + a[i] - a[0]
			ans = max(ans, cur)
		}
		if i+2 < n {
			cur := a[i+1] - a[i] + a[n-1] - a[i]
			ans = max(ans, cur)
		}
	}
	fmt.Fprintln(out, ans)
}

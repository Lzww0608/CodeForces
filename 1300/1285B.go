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

func check(x int) bool {
	for i := 2; i <= x/i; i++ {
		if x%i == 0 {
			return false
		}
	}

	return true
}

func solve(in *bufio.Reader, out *bufio.Writer) {
	var n int
	fmt.Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		fmt.Fscan(in, &a[i])
	}
	pre, suf := 0, 0
	for i := 0; i < n-1; i++ {
		pre += a[i]
		suf += a[n-i-1]
		if pre <= 0 || suf <= 0 {
			fmt.Fprintln(out, "NO")
			return
		}
	}

	fmt.Fprintln(out, "YES")
}

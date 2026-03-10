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
	b := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Fscan(in, &b[i])
	}

	left := make([]int, n)
	right := make([]int, n)
	left[0] = b[0]
	for i := 1; i < n; i++ {
		left[i] = max(left[i-1], b[i]+i)
	}
	right[n-1] = b[n-1] - (n - 1)
	for i := n - 2; i >= 0; i-- {
		right[i] = max(right[i+1], b[i]-i)
	}

	ans := 0
	for i := 1; i < n-1; i++ {
		ans = max(ans, b[i]+left[i-1]+right[i+1])
	}

	fmt.Fprintln(out, ans)
}

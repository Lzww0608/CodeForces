package main

import (
	"bufio"
	"fmt"
	"math"
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
	for i := range n {
		fmt.Fscan(in, &a[i])
	}

	ans := math.MinInt
	for i := max(0, n-k*2-1); i < n; i++ {
		for j := i + 1; j < n; j++ {
			ans = max(ans, (i+1)*(j+1)-k*(a[i]|a[j]))
		}
	}

	fmt.Fprintln(out, ans)
}

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
	sum, mx := 0, 0
	for i := range a {
		fmt.Fscan(in, &a[i])
		sum += a[i]
		mx = max(mx, a[i])
	}

	for i := n; i >= 1; i-- {
		need := max(mx*i, (sum+i-1)/i*i)
		if sum+k >= need {
			fmt.Fprintln(out, i)
			return
		}
	}
	fmt.Fprintln(out, 0)
}

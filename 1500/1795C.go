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
	b := make([]int, n+1)
	for i := range a {
		fmt.Fscan(in, &a[i])
	}
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &b[i])
		if i > 0 {
			b[i] += b[i-1]
		}
	}

	dif := make([]int, n+1)
	add := make([]int, n+1)
	cur := 0
	for i, x := range a {
		j := sort.SearchInts(b, b[i]+x+1) - 1
		dif[i]++
		dif[j]--
		cur += dif[i]
		add[j] += x - (b[j] - b[i])
		fmt.Fprintf(out, "%d ", cur*(b[i+1]-b[i])+add[i])
	}
	fmt.Fprintln(out)
}

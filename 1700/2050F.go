package main

import (
	"bufio"
	"fmt"
	"math/bits"
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
	var n, q int
	fmt.Fscan(in, &n, &q)
	a := make([]int, n)
	for i := range a {
		fmt.Fscan(in, &a[i])
		if i > 0 {
			a[i-1] = abs(a[i-1] - a[i])
		}
	}

	w := bits.Len(uint(n - 1))
	st := make([][]int, w+1)
	for i := range st {
		st[i] = make([]int, n-1)
	}
	st[0] = a[:n-1]

	for i := 1; i < w; i++ {
		for j := range n - 1 - (1 << i) + 1 {
			st[i][j] = gcd(st[i-1][j], st[i-1][j+(1<<(i-1))])
		}
	}

	var l, r int
	for range q {
		fmt.Fscan(in, &l, &r)
		l--
		r--
		if l == r {
			fmt.Fprintf(out, "0\n")
			continue
		}
		k := bits.Len(uint(r-l)) - 1
		fmt.Fprintf(out, "%d ", gcd(st[k][l], st[k][r-(1<<k)]))
	}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}

	return a
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

const N int = 1001

var moves [N]int

func init() {
	for i := 0; i < N; i++ {
		moves[i] = 12
	}
	moves[1] = 0

	for i := 1; i < N; i++ {
		for x := 1; x <= i; x++ {
			j := i + i/x
			if j < N && moves[j] > moves[i]+1 {
				moves[j] = moves[i] + 1
			}
		}
	}
}

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
	b := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &b[i])
	}
	c := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &c[i])
	}

	k = min(k, n*12)
	type pair struct {
		x, y int
	}
	a := []pair{}
	ans := 0
	for i := range n {
		if b[i] == 1 {
			ans += c[i]
		} else {
			a = append(a, pair{moves[b[i]], c[i]})
		}
	}

	f := make([]int, k+1)
	for _, v := range a {
		for x := k; x >= v.x; x-- {
			f[x] = max(f[x], f[x-v.x]+v.y)
		}
	}

	fmt.Fprintln(out, ans+slices.Max(f))
}

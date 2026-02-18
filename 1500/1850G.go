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
	var x, y int
	v := make(map[int]int)
	h := make(map[int]int)
	dg := make(map[int]int)
	udg := make(map[int]int)

	for range n {
		fmt.Fscan(in, &x, &y)
		v[x]++
		h[y]++
		dg[x-y]++
		udg[x+y]++
	}

	ans := 0
	for _, t := range v {
		ans += t * (t - 1)
	}
	for _, t := range h {
		ans += t * (t - 1)
	}
	for _, t := range dg {
		ans += t * (t - 1)
	}
	for _, t := range udg {
		ans += t * (t - 1)
	}

	fmt.Fprintln(out, ans)
}

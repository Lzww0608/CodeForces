package main

import (
	"bufio"
	"fmt"
	"os"
)

const N int = 1_000_000

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

	pre := make(map[int]int)
	mid := make(map[int]int)
	suf := make(map[int]int)
	m := make(map[int]int)

	ans := 0
	for i := 2; i < n; i++ {
		x, y, z := a[i], a[i-1], a[i-2]
		ans += pre[y*N+z] + mid[x*N+z] + suf[x*N+y]
		ans -= 3 * m[x*N*N+y*N+z]
		pre[y*N+z]++
		mid[x*N+z]++
		suf[x*N+y]++
		m[x*N*N+y*N+z]++
	}

	fmt.Fprintln(out, ans)
}

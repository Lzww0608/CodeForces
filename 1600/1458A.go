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

	var n, m int
	fmt.Fscan(in, &n, &m)
	a := make([]int, n)
	b := make([]int, m)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
	}
	for i := 0; i < m; i++ {
		fmt.Fscan(in, &b[i])
	}

	// a1 + bj, a2 + bj, ..., an + bj
	// a1 + bj, a2 - a1, ..., an - a1
	g := 0
	for i := 1; i < n; i++ {
		g = gcd(a[i]-a[0], g)
	}

	for i := 0; i < m; i++ {
		fmt.Fprintf(out, "%d ", abs(gcd(a[0]+b[i], g)))
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

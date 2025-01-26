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

	var n int
	fmt.Fscan(in, &n)
	a := make([]int, n)
	cnt := 0
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
		if a[i] == 1 {
			cnt++
		}
	}

	if cnt > 0 {
		fmt.Fprintln(out, n-cnt)
		return
	}
	g := a[0]
	for _, x := range a[1:] {
		g = gcd(g, x)
	}

	if g != 1 {
		fmt.Fprintln(out, -1)
		return
	}

	m := make([][]int, n)
	for i := range m {
		m[i] = make([]int, n)
		m[i][i] = a[i]
	}
	for d := 2; d <= n; d++ {
		for i := 0; i+d <= n; i++ {
			g = gcd(m[i][i+d-2], a[i+d-1])
			m[i][i+d-1] = g
			if g == 1 {
				fmt.Fprintln(out, n-1+d-1)
				return
			}
		}
	}
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

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

	var t, n int
	for fmt.Fscan(in, &t); t > 0; t-- {
		fmt.Fscan(in, &n)
		a := make([]int, n)
		for i := range a {
			fmt.Fscan(in, &a[i])
		}
		if n == 1 {
			fmt.Fprintln(out, 1)
			continue
		}
		sort.Ints(a)
		g := a[1] - a[0]
		for i := 2; i < n; i++ {
			g = gcd(g, a[i]-a[i-1])
		}
		b := a[n-1] + g
		for i := n - 1; i >= 1; i-- {
			if a[i]-a[i-1] > g {
				b = a[i] - g
				break
			}
		}
		mx := max(a[n-1], b)
		ans := (mx - b) / g
		for _, x := range a {
			ans += (mx - x) / g
		}

		fmt.Fprintln(out, ans)
	}
}

func gcd(x, y int) int {
	if x < 0 {
		x = -x
	}
	if y < 0 {
		y = -y
	}
	for y != 0 {
		x, y = y, x%y
	}

	return x
}

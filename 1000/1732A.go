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

	var t, n int
	for fmt.Fscan(in, &t); t > 0; t-- {
		fmt.Fscan(in, &n)
		a := make([]int, n)
		for i := range a {
			fmt.Fscan(in, &a[i])
		}
		ans := a[0]
		for i := 1; i < n; i++ {
			ans = gcd(ans, a[i])
		}
		if ans == 1 {
			fmt.Fprintln(out, 0)
			continue
		}
		if gcd(ans, n) == 1 {
			fmt.Fprintln(out, 1)
		} else if gcd(ans, n-1) == 1 {
			fmt.Fprintln(out, 2)
		} else {
			fmt.Fprintln(out, 3)
		}

	}
}

func gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}

	return x
}

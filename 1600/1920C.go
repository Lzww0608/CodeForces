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
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(in, &a[i])
	}

	ans := 0
	for k := 1; k*k <= n; k++ {
		if n%k != 0 {
			continue
		}
		b := []int{k}
		if k*k != n {
			b = append(b, n/k)
		}
		for _, p := range b {
			g := 0
			for i := range n {
				if i+p < n {
					g = gcd(abs(a[i]-a[i+p]), g)
					if g == 1 {
						break
					}
				} else {
					break
				}
			}
			if g != 1 {
				ans++
			}
		}
	}

	fmt.Fprintln(out, ans)
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}

	return a
}

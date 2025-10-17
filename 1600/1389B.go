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
	var n, k, z int
	fmt.Fscan(in, &n, &k, &z)
	a := make([]int, n)
	for i := range a {
		fmt.Fscan(in, &a[i])
	}

	ans, s, mx := 0, 0, 0
	for i := 0; i <= k; i++ {
		if i < n-1 {
			mx = max(mx, a[i]+a[i+1])
		}

		s += a[i]

		if i&1 == k&1 {
			tmp := (k - i) / 2
			if tmp <= z {
				ans = max(ans, s+mx*tmp)
			}
		}
	}

	fmt.Fprintln(out, ans)
}

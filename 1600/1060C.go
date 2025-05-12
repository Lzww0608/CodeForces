package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, m, k int
	fmt.Fscan(in, &n, &m)
	a := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i])
		a[i] += a[i-1]
	}
	b := make([]int, m+1)
	for i := 1; i <= m; i++ {
		fmt.Fscan(in, &b[i])
		b[i] += b[i-1]
	}

	fmt.Fscan(in, &k)

	c := make([]int, n+1)
	d := make([]int, m+1)
	for l := 1; l <= n; l++ {
		c[l] = math.MaxInt32
		for i := 0; i+l <= n; i++ {
			c[l] = min(c[l], a[i+l]-a[i])
		}
	}

	for l := 1; l <= m; l++ {
		d[l] = math.MaxInt32
		for i := 0; i+l <= m; i++ {
			d[l] = min(d[l], b[i+l]-b[i])
		}
	}

	ans := 0
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if k >= c[i]*d[j] {
				ans = max(ans, i*j)
			}
		}
	}

	fmt.Fprintln(out, ans)
}

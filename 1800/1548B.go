package main

import (
	"bufio"
	"fmt"
	"math/bits"
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
	for i := range a {
		fmt.Fscan(in, &a[i])
		if i > 0 {
			a[i-1] = abs(a[i-1] - a[i])
		}
	}

	if n == 1 {
		fmt.Fprintln(out, 1)
		return
	}

	n--
	w := bits.Len(uint(n))
	st := make([][]int, w)
	for i := range st {
		st[i] = make([]int, n)
	}
	st[0] = a

	for i := 1; i < w; i++ {
		for j := range n - (1 << i) + 1 {
			st[i][j] = gcd(st[i-1][j], st[i-1][j+(1<<(i-1))])
		}
	}

	ans := 1
	for i := 0; i < n && i+ans <= n; i++ {
		l, r := i-1, n
		for l+1 < r {
			mid := l + ((r - l) >> 1)
			if mid-i+1 <= ans {
				l = mid
			}

			k := bits.Len(uint(mid-i+1)) - 1
			g := gcd(st[k][i], st[k][mid+1-(1<<k)])
			if g == 1 {
				r = mid
			} else {
				ans = max(ans, mid-i+2)
				l = mid
			}
		}
	}

	fmt.Fprintln(out, ans)
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

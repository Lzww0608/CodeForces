package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

const MOD int = 1_000_000_007

var nums []string

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	for i := 1; i <= 1e18; i *= 2 {
		nums = append(nums, fmt.Sprintf("%d", i))
	}

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

	const mx = 2001
	f := make([][mx]int, n)
	for i := range f {
		for j := range f[i] {
			f[i][j] = math.MaxInt32
		}
	}
	f[0][a[0]] = a[0]
	for i := 1; i < n; i++ {
		x := a[i]
		for j := 0; j < mx; j++ {
			if j+x < mx {
				f[i][j+x] = min(f[i][j+x], max(f[i-1][j], j+x))
			}

			if j >= x {
				f[i][j-x] = min(f[i][j-x], f[i-1][j])
			} else {
				f[i][0] = min(f[i][0], f[i-1][j]+x-j)
			}
		}
	}

	ans := math.MaxInt32
	for j := range mx {
		ans = min(ans, f[n-1][j])
	}
	fmt.Fprintln(out, ans)
}

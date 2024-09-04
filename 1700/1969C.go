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

	var t, n, k int
	for fmt.Fscan(in, &t); t > 0; t-- {
		fmt.Fscan(in, &n, &k)
		a := make([]int, n)
		for i := range a {
			fmt.Fscan(in, &a[i])
		}

		f := make([][]int, n+1)
		for i := range f {
			f[i] = make([]int, k+1)
			for j := range f[i] {
				f[i][j] = math.MaxInt64 / 2
			}
		}

		f[0][0] = 0
		for i := 0; i < n; i++ {
			mn := a[i]
			for j := 0; i+j < n && j <= k; j++ {
				mn = min(mn, a[i+j])
				for x := 0; x+j <= k; x++ {
					f[i+j+1][x+j] = min(f[i+j+1][x+j], f[i][x]+mn*(j+1))
				}
			}
		}

		ans := f[n][0]
		for _, x := range f[n] {
			ans = min(ans, x)
		}
		fmt.Fprintln(out, ans)
	}
}

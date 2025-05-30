package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

const N int = 101

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)
	a := make([]int, n)
	b := make([]int, n)
	sum := 0
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i], &b[i])
		sum += b[i]
	}

	f := make([][]int, n*N)
	for i := range f {
		f[i] = make([]int, n+1)
		for j := range f[i] {
			f[i][j] = math.MinInt32
		}
	}
	f[0][0] = 0
	for i := 0; i < n; i++ {
		for j := n*N - 1; j >= a[i]; j-- {
			for k := n; k > 0; k-- {
				f[j][k] = max(f[j][k], f[j-a[i]][k-1]+b[i])
			}
		}
	}

	for i := 1; i <= n; i++ {
		ans := 0.0
		for j := 1; j < n*N; j++ {
			ans = max(ans, min(float64(j), float64(f[j][i])+float64(sum-f[j][i])*0.5))
		}
		fmt.Fprintf(out, "%f ", ans)
	}
}

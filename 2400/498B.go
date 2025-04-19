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

	var n, T int
	fmt.Fscan(in, &n, &T)
	p := make([]float64, n+1)
	t := make([]int, n+1)

	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &p[i], &t[i])
		p[i] /= 100.0
	}

	var ans float64
	f := make([][]float64, n+1)
	for i := range f {
		f[i] = make([]float64, T+1)
	}
	f[0][0] = 1.0
	for i := 1; i <= n; i++ {
		pow := math.Pow(float64(1.0-p[i]), float64(t[i]-1))
		var sum, tmp float64
		for j := 1; j <= T; j++ {
			sum = sum*(1.0-p[i]) + f[i-1][j-1]*p[i]
			if j >= t[i] {
				sum -= f[i-1][j-t[i]] * p[i] * pow
				tmp = sum + f[i-1][j-t[i]]*pow
			} else {
				tmp = sum
			}
			f[i][j] += tmp
			ans += f[i][j]
		}
	}

	fmt.Fprintln(out, ans)
}

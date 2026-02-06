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

	var n int
	fmt.Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		fmt.Fscan(in, &a[i])
	}

	f := make([][]int, n+1)
	for i := range f {
		f[i] = make([]int, n+1)
		f[i][i] = 1
	}
	for i := n - 2; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			cur := math.MaxInt32
			if a[i] == a[j] {
				if i+1 == j {
					cur = 1
				} else {
					cur = f[i+1][j-1]
				}
			} else {
				cur = min(f[i+1][j], f[i][j-1]) + 1
			}

			for k := i; k < j; k++ {
				cur = min(cur, f[i][k]+f[k+1][j])
			}
			f[i][j] = cur
		}
	}

	fmt.Fprintln(out, f[0][n-1])
}

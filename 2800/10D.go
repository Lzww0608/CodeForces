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

	var n, m int
	fmt.Fscan(in, &n)
	a := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i])
	}
	fmt.Fscan(in, &m)
	b := make([]int, m+1)
	for i := 1; i <= m; i++ {
		fmt.Fscan(in, &b[i])
	}

	f := make([][]int, n+1)
	f[0] = make([]int, m+1)
	ans, l, r := 0, 0, 0
	for i := 1; i <= n; i++ {
		f[i] = make([]int, m+1)
		mx := 0
		for j := 1; j <= m; j++ {
			if a[i] != b[j] {
				f[i][j] = f[i-1][j]
			}
			if a[i] > b[j] && mx < f[i-1][j] {
				mx = f[i-1][j]
			}

			if a[i] == b[j] {
				f[i][j] = mx + 1
				if f[i][j] > ans {
					ans = f[i][j]
					l, r = i, j
				}
			}
		}
	}

	fmt.Fprintln(out, ans)

	var print func(int, int)
	print = func(l, r int) {
		if f[l][r] == 1 {
			fmt.Fprintf(out, "%d ", b[r])
			return
		}

		if l == 0 || r == 0 {
			return
		}
		if f[l][r] == f[l-1][r] {
			print(l-1, r)
			return
		}

		for i := r - 1; i >= 1; i-- {
			if b[i] < b[r] && f[l][r] == f[l-1][i]+1 {
				print(l-1, i)
				fmt.Fprintf(out, "%d ", b[r])
				return
			}
		}
	}
	if ans != 0 {
		print(l, r)
	}

}

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

	var n int
	fmt.Fscan(in, &n)
	a := make([][]int, n)
	row := make([]int, n)
	col := make([]int, n)
	dg, udg := 0, 0
	x, y := 0, 0
	for i := range a {
		a[i] = make([]int, n)
		for j := range a[i] {
			fmt.Fscan(in, &a[i][j])
			if a[i][j] == 0 {
				x, y = i, j
			}

			row[i] += a[i][j]
			col[j] += a[i][j]
			if i == j {
				dg += a[i][j]
			}
			if i == n-j-1 {
				udg += a[i][j]
			}
		}
	}

	if n == 1 {
		fmt.Fprintln(out, 1)
		return
	}

	k := 0
	if x > 0 {
		k = row[x-1] - row[x]
	} else {
		k = row[x+1] - row[x]
	}

	if k <= 0 {
		fmt.Fprintln(out, -1)
		return
	}

	row[x] += k
	col[y] += k
	if x == y {
		dg += k
	}
	if x == n-y-1 {
		udg += k
	}
	//fmt.Fprintln(out, k, col, row, dg, udg)
	for i := 1; i < n; i++ {
		if row[i-1] != row[i] || col[i-1] != col[i] || row[i] != col[i] {
			fmt.Fprintln(out, -1)
			return
		}
	}

	if dg != udg || dg != row[0] {
		fmt.Fprintln(out, -1)
		return
	}
	fmt.Fprintln(out, k)
}

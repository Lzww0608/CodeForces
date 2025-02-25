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
	a := make([]int, n)
	vis := make([]bool, n)
	cntZero, cntNeg := 0, 0
	pos := -1
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
		if a[i] == 0 {
			cntZero++
			vis[i] = true
		} else if a[i] < 0 {
			cntNeg++
			if pos == -1 || abs(a[i]) < abs(a[pos]) {
				pos = i
			}
		}
	}

	if cntNeg&1 == 1 {
		vis[pos] = true
	}

	if cntZero == n || (cntZero == n-1 && cntNeg == 1) {
		for i := 0; i < n-1; i++ {
			fmt.Fprintln(out, 1, i+1, i+2)
		}
		return
	}

	j := -1
	for i := 0; i < n; i++ {
		if vis[i] {
			if j != -1 {
				fmt.Fprintln(out, 1, j+1, i+1)
			}
			j = i
		}
	}

	if j != -1 {
		fmt.Fprintln(out, 2, j+1)
	}

	j = -1
	for i := 0; i < n; i++ {
		if !vis[i] {
			if j != -1 {
				fmt.Fprintln(out, 1, j+1, i+1)
			}
			j = i
		}
	}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

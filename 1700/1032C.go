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
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
	}
	f := make([][5]int, n)
	for i := 1; i < n; i++ {
		for j := 0; j < 5; j++ {
			f[i][j] = -1
		}
	}
	for i := 0; i < 5; i++ {
		f[0][i] = i
	}

	for i := 1; i < n; i++ {
		for j := 0; j < 5; j++ {
			if f[i-1][j] == -1 {
				continue
			}
			for k := 0; k < 5; k++ {
				if j == k {
					continue
				}
				if a[i] > a[i-1] {
					if k > j {
						f[i][k] = j
					}
				} else if a[i] < a[i-1] {
					if k < j {
						f[i][k] = j
					}
				} else {
					if k != j {
						f[i][k] = j
					}
				}

			}
		}

	}

	ans := make([]int, n)
	t := -1
	for i := 0; i < 5; i++ {
		if f[n-1][i] != -1 {
			t = i
			break
		}
	}
	if t == -1 {
		fmt.Fprintln(out, -1)
		return
	}
	for i := n - 1; i >= 0; i-- {
		ans[i] = t + 1
		t = f[i][t]
	}
	for _, v := range ans {
		fmt.Fprint(out, v, " ")
	}
}

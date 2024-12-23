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
	fmt.Fscan(in, &n, &m)
	a := make([][]int, n)
	xor := 0
	ans := make([]int, n)
	for i := range a {
		a[i] = make([]int, m)
		for j := range a[i] {
			fmt.Fscan(in, &a[i][j])
		}
	}

	for i := range ans {
		ans[i] = 1
		xor ^= a[i][0]
	}

	if xor != 0 {
		fmt.Fprintln(out, "TAK")
		for _, x := range ans {
			fmt.Fprint(out, x, " ")
		}
		return
	}
	for i := 0; i < n; i++ {
		for j := 1; j < m; j++ {
			if a[i][j] != a[i][0] {
				ans[i] = j + 1
				fmt.Fprintln(out, "TAK")
				for _, x := range ans {
					fmt.Fprint(out, x, " ")
				}
				return
			}
		}
	}

	fmt.Fprintln(out, "NIE")
}

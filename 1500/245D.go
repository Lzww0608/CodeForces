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
	for i := 0; i < n; i++ {
		a[i] = make([]int, n)
		for j := 0; j < n; j++ {
			fmt.Fscan(in, &a[i][j])
		}
	}

	ans := make([]int, n)
	for i := 0; i < n; i++ {
		cur := 0
		for j := 0; j < n; j++ {
			if i == j {
				continue
			}
			cur |= a[i][j]
		}
		ans[i] = cur
	}
	for _, x := range ans {
		fmt.Fprint(out, x, " ")
	}
}

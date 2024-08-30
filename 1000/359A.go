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

	var m, n int
	fmt.Fscan(in, &m, &n)
	a := make([][]int, m)
	for i := range a {
		a[i] = make([]int, n)
		for j := range a[i] {
			fmt.Fscan(in, &a[i][j])
		}
	}

	for i, row := range a {
		for j, x := range row {
			if x == 1 && (i == 0 || i == m-1 || j == 0 || j == n-1) {
				fmt.Fprintln(out, 2)
				return
			}
		}
	}

	fmt.Fprintln(out, 4)
}

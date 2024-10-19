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

	var n, m, k int
	fmt.Fscan(in, &n, &m, &k)
	b := make([][]int, n)
	for i := range b {
		b[i] = make([]int, m)
	}
	row := make([][2]int, n)
	col := make([][2]int, m)
	var x, y, z int
	for i := 1; i <= k; i++ {
		fmt.Fscan(in, &x, &y, &z)
		if x == 1 {
			row[y-1] = [2]int{i, z}
		} else {
			col[y-1] = [2]int{i, z}
		}

	}
	for i, v := range b {
		for j := range v {
			x, y = row[i][0], col[j][0]
			if x > y {
				fmt.Fprint(out, row[i][1], " ")
			} else {
				fmt.Fprint(out, col[j][1], " ")
			}
		}
		fmt.Fprintln(out)
	}
}

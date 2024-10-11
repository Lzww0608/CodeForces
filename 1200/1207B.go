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
	b := make([][]int, n)
	for i := range a {
		a[i] = make([]int, m)
		b[i] = make([]int, m)
		for j := range a[i] {
			fmt.Fscan(in, &a[i][j])
		}
	}

	ans := [][]int{}
	for i := 0; i < n-1; i++ {
		for j := 0; j < m-1; j++ {
			if a[i][j] == 1 && a[i+1][j] == 1 && a[i][j+1] == 1 && a[i+1][j+1] == 1 {
				b[i][j], b[i+1][j], b[i][j+1], b[i+1][j+1] = 1, 1, 1, 1
				ans = append(ans, []int{i + 1, j + 1})
			}
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if a[i][j] != b[i][j] {
				fmt.Fprintln(out, -1)
				return
			}
		}
	}
	fmt.Fprintln(out, len(ans))
	for _, v := range ans {
		fmt.Fprintln(out, v[0], v[1])
	}

}

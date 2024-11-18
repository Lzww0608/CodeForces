package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var t int
	for fmt.Fscan(in, &t); t > 0; t-- {
		solve(in, out)
	}
}

func solve(in *bufio.Reader, out *bufio.Writer) {
	var n, m int
	fmt.Fscan(in, &n, &m)
	a := make([][]int, n)
	for i := range a {
		a[i] = make([]int, m)
		for j := 0; j < m; j++ {
			fmt.Fscan(in, &a[i][j])
		}
	}

	pos := []int{}
	for i := 0; i < n && len(pos) == 0; i++ {
		b := make([]int, m)
		copy(b, a[i])
		sort.Ints(b)
		for j := 0; j < m; j++ {
			if a[i][j] != b[j] {
				pos = append(pos, j)
			}
		}
	}

	if len(pos) == 0 {
		fmt.Fprintln(out, 1, 1)
		return
	} else if len(pos) > 2 {
		fmt.Fprintln(out, -1)
		return
	}

	for i := 0; i < n; i++ {
		a[i][pos[0]], a[i][pos[1]] = a[i][pos[1]], a[i][pos[0]]
	}

	for i := 0; i < n; i++ {
		for j := 1; j < m; j++ {
			if a[i][j] < a[i][j-1] {
				fmt.Fprintln(out, -1)
				return
			}
		}
	}

	fmt.Fprintln(out, pos[0]+1, pos[1]+1)
}

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
	b := make([][]int, n)
	id := make([]int, m*n)
	for i := range id {
		id[i] = i
	}

	for i := range a {
		b[i] = make([]int, m)
		a[i] = make([]int, m)
		for j := range a[i] {
			fmt.Fscan(in, &a[i][j])
		}
	}

	sort.Slice(id, func(i, j int) bool {
		x, y := id[i]/m, id[i]%m
		xx, yy := id[j]/m, id[j]%m
		return a[x][y] < a[xx][yy]
	})
	//fmt.Fprintln(out, id)
	for j, i := range id[:m] {
		x, y := i/m, i%m
		b[x][j] = a[x][y]
	}

	p := make([]int, n)
	for _, i := range id[m:] {
		x, y := i/m, i%m
		for b[x][p[x]] != 0 {
			p[x]++
		}
		b[x][p[x]] = a[x][y]
	}

	for i := range b {
		for _, x := range b[i] {
			fmt.Fprint(out, x, " ")
		}
		fmt.Fprintln(out)
	}
}

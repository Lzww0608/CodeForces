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

	var t int
	for fmt.Fscan(in, &t); t > 0; t-- {
		solve(in, out)
	}
}

func solve(in *bufio.Reader, out *bufio.Writer) {
	var n int
	fmt.Fscan(in, &n)
	a := make([][]int, n)
	for i := range a {
		a[i] = make([]int, n)
		for j := range a[i] {
			fmt.Fscan(in, &a[i][j])
		}
	}
	b := make([]int, n)
	for i := range a {
		cur := (1 << 30) - 1
		for j := range a[i] {
			if j != i {
				cur &= a[i][j]
			}
		}
		b[i] = cur
	}

	for i := range a {
		for j := range a[i] {
			if i != j && b[i]|b[j] != a[i][j] {
				fmt.Fprintln(out, "NO")
				return
			}
		}
	}
	fmt.Fprintln(out, "YES")
	for _, x := range b {
		fmt.Fprint(out, x, " ")
	}
	fmt.Fprintln(out)
}

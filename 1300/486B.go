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
	b := make([][]int, m)
	row := make([]bool, m)
	col := make([]bool, n)
	for i := range a {
		a[i] = make([]int, n)
		b[i] = make([]int, n)
		for j := range a[i] {
			a[i][j] = 1
			fmt.Fscan(in, &b[i][j])
			if b[i][j] == 0 {
				row[i] = true
				col[j] = true
			}
		}
	}

	for i := range m {
		for j := range n {
			if row[i] || col[j] {
				a[i][j] = 0
			}
		}
	}

	for i := range m {
		for j := range n {
			if b[i][j] == 1 {
				f := false
				for k := range m {
					if a[k][j] == 1 {
						f = true
						break
					}
				}
				for k := range n {
					if a[i][k] == 1 {
						f = true
						break
					}
				}
				if !f {
					fmt.Fprintln(out, "NO")
					return
				}
			}
		}
	}

	fmt.Fprintln(out, "YES")
	for i := range m {
		for j := range n {
			fmt.Fprintf(out, "%d ", a[i][j])
		}
		fmt.Fprintln(out)
	}
}

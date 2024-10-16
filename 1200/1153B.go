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

	var n, m, h int
	fmt.Fscan(in, &n, &m, &h)
	a := make([]int, n)
	b := make([]int, m)
	c := make([][]int, n)
	for i := range b {
		fmt.Fscan(in, &b[i])
	}
	for i := range a {
		fmt.Fscan(in, &a[i])
	}
	for i := range c {
		c[i] = make([]int, m)
		for j := range c[i] {
			fmt.Fscan(in, &c[i][j])
			if c[i][j] == 1 {
				c[i][j] = min(a[i], b[j])
			}
		}
	}

	for i := range c {
		for _, x := range c[i] {
			fmt.Fprint(out, x, " ")
		}
		fmt.Fprintln(out)
	}
}

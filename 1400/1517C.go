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

	solve(in, out)
}

func solve(in *bufio.Reader, out *bufio.Writer) {
	var n int
	fmt.Fscan(in, &n)
	a := make([][]int, n)
	for i := range a {
		a[i] = make([]int, i+1)
		fmt.Fscan(in, &a[i][i])
	}

	for i := 0; i < n; i++ {
		cnt := a[i][i]
		x, y := i, i
		for cnt > 0 {
			a[x][y] = a[i][i]
			if y-1 >= 0 && a[x][y-1] == 0 {
				y = y - 1
			} else {
				x = x + 1
			}
			cnt--
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j <= i; j++ {
			fmt.Fprint(out, a[i][j], " ")
		}
		fmt.Fprintln(out)
	}
}

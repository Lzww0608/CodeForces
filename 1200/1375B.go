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
	var n, m int
	fmt.Fscan(in, &n, &m)
	a := make([][]int, n)
	for i := range a {
		a[i] = make([]int, m)
		for j := range a[i] {
			fmt.Fscan(in, &a[i][j])
		}
	}

	if a[0][0] >= 3 || a[n-1][0] >= 3 || a[0][m-1] >= 3 || a[n-1][m-1] >= 3 {
		fmt.Fprintln(out, "NO")
		return
	}

	for i := 1; i < n-1; i++ {
		if a[i][0] >= 4 || a[i][m-1] >= 4 {
			fmt.Fprintln(out, "NO")
			return
		}
	}

	for j := 1; j < m-1; j++ {
		if a[0][j] >= 4 || a[n-1][j] >= 4 {
			fmt.Fprintln(out, "NO")
			return
		}
	}

	for i := 1; i < n-1; i++ {
		for j := 1; j < m-1; j++ {
			if a[i][j] > 4 {
				fmt.Fprintln(out, "NO")
				return
			}
		}

	}

	fmt.Fprintln(out, "YES")
	for i := 0; i < n; i++ {
		if i == 0 || i == n-1 {
			fmt.Fprint(out, 2, " ")
		} else {
			fmt.Fprint(out, 3, " ")
		}

		for j := 1; j < m-1; j++ {
			if i == 0 || i == n-1 {
				fmt.Fprint(out, 3, " ")
			} else {
				fmt.Fprint(out, 4, " ")
			}
		}

		if i == 0 || i == n-1 {
			fmt.Fprintln(out, 2, " ")
		} else {
			fmt.Fprintln(out, 3, " ")
		}
	}
}

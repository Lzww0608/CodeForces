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
	if m == 1 {
		fmt.Fprintln(out, 0)
	} else if n >= m {
		fmt.Fprintln(out, m)
	} else {
		fmt.Fprintln(out, n+1)
	}

	for i := 0; i < min(n, m-1); i++ {
		for j := 0; j < m; j++ {
			fmt.Fprint(out, (i+j)%m, " ")
		}
		fmt.Fprintln(out)
	}
	for i := m - 1; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Fprint(out, j, " ")
		}
		fmt.Fprintln(out)
	}
}

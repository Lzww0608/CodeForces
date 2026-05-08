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

	var q int
	for fmt.Fscan(in, &q); q > 0; q-- {
		solve(in, out)
	}
}

func solve(in *bufio.Reader, out *bufio.Writer) {
	var n int
	fmt.Fscan(in, &n)
	var s [2]string
	fmt.Fscan(in, &s[0], &s[1])

	i, j := 0, 0
	for j < n {
		if s[i][j] >= '3' {
			if s[i^1][j] < '3' {
				fmt.Fprintln(out, "NO")
				return
			}
			i ^= 1
		}
		j++
	}

	if i == 1 && j == n {
		fmt.Fprintln(out, "YES")
	} else {
		fmt.Fprintln(out, "NO")
	}
}

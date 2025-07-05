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
	s := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &s[i])
	}

	for i := n - 2; i >= 0; i-- {
		for j := n - 2; j >= 0; j-- {
			if s[i][j] == '1' && s[i+1][j] != '1' && s[i][j+1] != '1' {
				fmt.Fprintln(out, "NO")
				return
			}
		}
	}

	fmt.Fprintln(out, "YES")
}

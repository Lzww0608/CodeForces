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
	var s string
	fmt.Fscan(in, &s)

	for i := 0; i < n; i++ {
		if s[i] == '0' {
			if i >= n/2 {
				fmt.Fprintln(out, 1, i+1)
				fmt.Fprintln(out, 1, i)
			} else {
				fmt.Fprintln(out, i+1, n)
				fmt.Fprintln(out, i+2, n)
			}
			return
		}
	}

	fmt.Fprintln(out, 1, n-1)
	fmt.Fprintln(out, 2, n)
}

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
	if n == 2 {
		fmt.Fprintln(out, 1, 1)
		fmt.Fprintln(out, 1, 2)
		return
	}
	fmt.Fprintln(out, n, n-1)
	fmt.Fprintln(out, n, n)
	for i := 1; i < n-1; i++ {
		fmt.Fprintln(out, i, i)
	}
}

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
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i * (i*2 - 1)
	}
	fmt.Fprintln(out, sum, n*2)
	for i := n; i >= 1; i-- {
		fmt.Fprint(out, 1, i, " ")
		for j := 1; j <= n; j++ {
			fmt.Fprint(out, j, " ")
		}
		fmt.Fprintln(out)
		fmt.Fprint(out, 2, i, " ")
		for j := 1; j <= n; j++ {
			fmt.Fprint(out, j, " ")
		}
		fmt.Fprintln(out)
	}
}

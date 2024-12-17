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
	if n <= 3 {
		fmt.Fprintln(out, -1)
		return
	}
	if n%4 == 0 {
		output(out, n)
		fmt.Fprintln(out)
		return
	}
	t := n % 4
	output(out, n-4-t)
	k := n - 4 - t + 1
	if t == 1 {
		fmt.Fprintln(out, k+1, k+3, k, k+2, k+4)
	} else if t == 2 {
		fmt.Fprint(out, k, k+2, k+4, k+1, k+3, k+5)
	} else if t == 3 {
		fmt.Fprint(out, k+2, k+5, k+1, k+3, k+6, k+4, k)
	}
	fmt.Fprintln(out)
	return
}

func output(out *bufio.Writer, n int) {
	for i := 1; i <= n; i += 4 {
		fmt.Fprint(out, i+1, i+3, i, i+2, " ")
	}
	return
}

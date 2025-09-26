package main

import (
	"bufio"
	"fmt"
	"io"
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

func solve(in io.Reader, out io.Writer) {
	var n, x, y int
	fmt.Fscan(in, &n, &x, &y)
	a := make([]int, n)
	for i := range a {
		fmt.Fscan(in, &a[i])
		x ^= a[i]
	}

	if x&1 == y&1 {
		fmt.Fprintln(out, "Alice")
	} else {
		fmt.Fprintln(out, "Bob")
	}
}

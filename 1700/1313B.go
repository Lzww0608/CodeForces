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
	var n, x, y int
	fmt.Fscan(in, &n, &x, &y)
	mn := max(1, min(n, x+y-n+1))
	mx := min(n, max(1, x+y-1))
	fmt.Fprintln(out, mn, mx)
}

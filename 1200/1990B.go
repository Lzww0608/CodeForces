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
	for i := 1; i <= n; i++ {
		if (i < y && (y-i)&1 == 1) || (i > x && (x-i)&1 == 1) {
			fmt.Fprint(out, -1, " ")
		} else {
			fmt.Fprint(out, 1, " ")
		}
	}
	fmt.Fprintln(out)
}

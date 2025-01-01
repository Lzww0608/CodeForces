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
	var x1, y1, x2, y2 int
	fmt.Fscan(in, &x1, &y1, &x2, &y2)
	if x1 == x2 || y1 == y2 {
		fmt.Fprintln(out, abs(x1-x2)+abs(y1-y2))
	} else {
		fmt.Fprintln(out, abs(x1-x2)+abs(y1-y2)+2)
	}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

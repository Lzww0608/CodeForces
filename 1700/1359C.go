package main

import (
	"bufio"
	"fmt"
	"math"
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
	var h, c, t int
	fmt.Fscan(in, &h, &c, &t)

	if h+c-t*2 >= 0 {
		fmt.Fprintln(out, 2)
		return
	} else if t >= h {
		fmt.Fprintln(out, 1)
		return
	} 

	n := ((h - c) / (2 * t - c - h) - 1) / 2
	if (4 * n * n + 6 * n + 1) * c + (4 * n * n + 10 * n  + 5) * h <= 2 * (2 * n + 1) * (2 * n + 3) * t {
	fmt.Fprintln(out, 2 * n + 1)
} else {
	fmt.Fprintln(out, 2 * n + 3)
}
}

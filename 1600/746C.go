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

	var s, x1, x2 int
	fmt.Fscan(in, &s, &x1, &x2)
	var t1, t2 int
	fmt.Fscan(in, &t1, &t2)
	var p, d int
	fmt.Fscan(in, &p, &d)

	a := abs(x1-x2) * t2
	b := abs(p-x2) * t1
	if d == 1 {
		if x1 < x2 {
			if p > x1 {
				b = (s + s + x2 - p) * t1
			}
		} else {
			b = (s - p + s - x2) * t1
		}
	} else {
		if x1 < x2 {
			b = (p + x2) * t1
		} else {
			if p < x1 {
				b = (s * 2 + p - x2) * t1
			}
		}
	}

	fmt.Fprintln(out, min(a, b))
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

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

	var n, x, y int
	fmt.Fscan(in, &n)

	xa, ya := math.MaxInt32, math.MaxInt32
	xb, yb := math.MinInt32, math.MinInt32
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &x, &y)
		if x < xa || x == xa && y < ya {
			xa, ya = x, y
		}
	}
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &x, &y)
		if x > xb || x == xb && y > yb {
			xb, yb = x, y
		}
	}
	fmt.Fprintln(out, xa+xb, ya+yb)
}

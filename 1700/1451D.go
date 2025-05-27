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
	var d, k float64
	fmt.Fscan(in, &d, &k)
	// dis(xk, xk) <= d
	// 2 * x * x * k * k == d * d
	// sqrt(2) * x * k == d
	t := float64(int(d / math.Sqrt(2.0) / k))
	x := (t + 1) * k
	y := t * k
	if x*x+y*y > d*d {
		fmt.Fprintln(out, "Utkarsh")
	} else {
		fmt.Fprintln(out, "Ashish")
	}
}

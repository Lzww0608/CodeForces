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
	var n int
	fmt.Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		fmt.Fscan(in, &a[i])
	}

	ans := 0
	// x <= y
	x, y := math.MaxInt32, math.MaxInt32
	for _, v := range a {
		if v > y {
			x, y = y, v
			ans++
		} else if v <= x {
			x = v
		} else {
			y = v
		}
	}
	fmt.Fprintln(out, ans)
}

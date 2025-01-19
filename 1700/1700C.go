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
	a := make([]int, n)
	for i := range a {
		fmt.Fscan(in, &a[i])
	}
	// 1. a[1] -= 1, a[i+1] += 1
	// 2. a[i] -= 1
	// 3. a[1] += 1
	x := a[0]
	ans := 0
	for i := 0; i < n-1; i++ {
		y := a[i+1] - a[i]
		ans += abs(y)
		if y < 0 {
			// 1
			x += y
		}
	}
	ans += abs(x)
	fmt.Fprintln(out, ans)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

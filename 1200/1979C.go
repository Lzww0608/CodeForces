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
	lcm := 1
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
		lcm = lcm * a[i] / gcd(lcm, a[i])
	}
	sum := 0
	for i := range a {
		sum += lcm / a[i]
	}
	if sum >= lcm {
		fmt.Fprintln(out, -1)
		return
	}
	for i := 0; i < n; i++ {
		fmt.Fprint(out, lcm/a[i], " ")
	}
	fmt.Fprintln(out)

}

func gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}

	return x
}

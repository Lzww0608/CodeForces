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
	b := make([]int, n)
	for i := range a {
		fmt.Fscan(in, &a[i])
	}
	for i := 0; i < n-3; i += 2 {
		x, y := abs(a[i]), abs(a[i+1])
		g := gcd(x, y)
		if a[i] < 0 && a[i+1] > 0 || a[i] > 0 && a[i+1] < 0 {
			b[i], b[i+1] = y/g, x/g
		} else {
			b[i], b[i+1] = -y/g, x/g
		}
	}

	if n&1 == 0 {
		i := n - 2
		x, y := abs(a[i]), abs(a[i+1])
		g := gcd(x, y)
		if a[i] < 0 && a[i+1] > 0 || a[i] > 0 && a[i+1] < 0 {
			b[i], b[i+1] = y/g, x/g
		} else {
			b[i], b[i+1] = -y/g, x/g
		}
	} else {
		i := n - 3
		if a[i]+a[i+1] != 0 {
			b[i], b[i+1], b[i+2] = -a[i+2], -a[i+2], a[i]+a[i+1]
		} else if a[i+1]+a[i+2] != 0 {
			b[i], b[i+1], b[i+2] = a[i+2]+a[i+1], -a[i], -a[i]
		} else {
			b[i], b[i+1], b[i+2] = -a[i+1], a[i]+a[i+2], -a[i+1]
		}
	}

	for _, x := range b {
		fmt.Fprint(out, x, " ")
	}
	fmt.Fprintln(out)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

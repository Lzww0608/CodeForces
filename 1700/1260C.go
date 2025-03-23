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
	var a, b, k int
	fmt.Fscan(in, &a, &b, &k)
	a, b = min(a, b), max(a, b)

	t := (b - 1) / a
	g := gcd(a, b)
	if (b-1)%a >= g {
		t++
	}
	if t >= k {
		fmt.Fprintln(out, "REBEL")
	} else {
		fmt.Fprintln(out, "OBEY")
	}
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

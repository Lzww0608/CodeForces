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
	var n, x int
	fmt.Fscan(in, &n, &x)

	ans := 0
	for a := 1; a < x; a++ {
		for b := 1; a+b < x && a*b < n; b++ {
			y := min((n-a*b)/(a+b), x-a-b)
			ans += y
		}
	}

	fmt.Fprintln(out, ans)
}

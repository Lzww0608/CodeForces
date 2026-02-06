package main

import (
	"bufio"
	"fmt"
	"io"
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

func solve(in io.Reader, out io.Writer) {
	var n int
	fmt.Fscan(in, &n)
	var s string
	fmt.Fscan(in, &s)

	type pair struct {
		x, y int
	}
	cnt := make(map[pair]int)

	x, y := 0, 0
	for _, c := range s {
		if c == 'D' {
			x++
		} else {
			y++
		}
		g := gcd(x, y)
		cnt[pair{x / g, y / g}]++
		fmt.Fprintf(out, "%d ", cnt[pair{x / g, y / g}])
	}
	fmt.Fprintln(out)
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}

	return a
}

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
	var n, s int
	fmt.Fscan(in, &n, &s)

	calc := func(x int) (cnt int) {
		for x > 0 {
			cnt += x % 10
			x /= 10
		}

		return
	}

	ans := 0
	for t := 1; calc(n) > s; t *= 10 {
		x := n / t % 10
		y := (10 - x) % 10
		ans += y * t
		n += y * t
	}

	fmt.Fprintln(out, ans)
}

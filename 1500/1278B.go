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
	var a, b int
	fmt.Fscan(in, &a, &b)
	if a > b {
		a, b = b, a
	}
	d := b - a
	ans := 0
	sum := 0

	for sum < d {
		ans++
		sum += ans
	}

	if (sum-d)&1 == 1 {
		ans++
	}
	fmt.Fprintln(out, ans)
}

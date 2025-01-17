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
	var a, b, c int
	fmt.Fscan(in, &a, &b, &c)
	sum := a + b + c
	if sum%9 == 0 {
		k := sum / 9
		if a >= k && b >= k && c >= k {
			fmt.Fprintln(out, "YES")
			return
		}
	}
	fmt.Fprintln(out, "NO")
}

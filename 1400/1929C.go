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
	var k, x, a int
	fmt.Fscan(in, &k, &x, &a)
	if a <= x {
		fmt.Fprintln(out, "NO")
		return
	}
	sum := 0
	for i := 0; i <= x; i++ {
		t := sum/(k-1) + 1
		sum += t
		if sum > a {
			fmt.Fprintln(out, "NO")
			return
		}
	}
	fmt.Fprintln(out, "YES")
}

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
	xor := 0
	for i := 1; i <= n-3; i++ {
		fmt.Fprint(out, i, " ")
		xor ^= i
	}
	xor ^= (1 << 29)
	xor ^= (1 << 30)
	fmt.Fprintln(out, 1<<29, 1<<30, xor)
}

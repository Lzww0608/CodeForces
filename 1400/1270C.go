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
	fmt.Fscan(in, &n)
	sum, xor := 0, 0
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &x)
		sum += x
		xor ^= x
	}

	if sum == xor*2 {
		fmt.Fprintln(out, 0)
	} else {
		fmt.Fprintln(out, 2)
		fmt.Fprintln(out, xor, sum+xor)
	}
}

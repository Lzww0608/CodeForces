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
	if a == b {
		fmt.Fprintln(out, -1)
		return
	}

	k := (1 << 50) - max(a, b)
	fmt.Fprintln(out, k)
}

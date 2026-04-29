package main

import (
	"bufio"
	"fmt"
	"math/bits"
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

	if n&1 == 1 {
		fmt.Fprintln(out, "Bob")
	} else {
		if n&(n-1) == 0 && bits.Len(uint(n))&1 == 0 {
			fmt.Fprintln(out, "Bob")
		} else {
			fmt.Fprintln(out, "Alice")
		}
	}
}

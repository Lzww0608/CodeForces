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
	var l, r int
	fmt.Fscan(in, &l, &r)
	var a, b, c int
	k := bits.Len(uint(l^r)) - 1
	a = l | ((1 << k) - 1)
	b = a + 1
	if b != r {
		c = r
	} else {
		c = l
	}

	fmt.Fprintln(out, a, b, c)
}

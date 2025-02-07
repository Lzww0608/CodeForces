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

	var a, b uint64
	fmt.Fscan(in, &a, &b)

	if a < b {
		fmt.Fprintln(out, -1)
		return
	}

	z := a - b
	if z&1 == 1 {
		fmt.Fprintln(out, -1)
		return
	}

	z >>= 1
	x := z
	y := b + z

	if x+y == a {
		fmt.Fprintln(out, x, y)
	} else {
		fmt.Fprintln(out, -1)
	}
}

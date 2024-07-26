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

	var t, a, b int
	for fmt.Fscan(in, &t); t > 0; t-- {
		fmt.Fscan(in, &a, &b)
		x := gcd(a, b)
		if b%a == 0 {
			fmt.Fprintln(out, b/a*b)
		} else {
			fmt.Fprintln(out, a*b/x)
		}
	}
}

func gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}

	return x
}

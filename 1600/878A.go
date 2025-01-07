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

	var n, t int
	fmt.Fscan(in, &n)
	var s string

	x, y := 0, 1023
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &s, &t)
		if s[0] == '|' {
			x |= t
			y |= t
		} else if s[0] == '&' {
			x &= t
			y &= t
		} else if s[0] == '^' {
			x ^= t
			y ^= t
		}
	}
	//fmt.Fprintln(out, x, y)
	a, b, c := 1023, 0, 0
	for i := 0; i < 10; i++ {
		k := 1 << i
		if x&k == 0 {
			if y&k == 0 {
				a -= k
			} else {

			}
		} else {
			if y&k == 0 {
				c |= k
			} else {
				b |= k
			}
		}
	}
	fmt.Fprintln(out, 3)
	fmt.Fprintf(out, "& %d\n", a)
	fmt.Fprintf(out, "| %d\n", b)
	fmt.Fprintf(out, "^ %d\n", c)
}

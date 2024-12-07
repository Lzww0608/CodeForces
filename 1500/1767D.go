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

	var n int
	fmt.Fscan(in, &n)
	l, r := 1, 1<<n
	a, b := 1, 1
	var s string
	fmt.Fscan(in, &s)
	for i := range s {
		if s[i] == '1' {
			l += a
			a *= 2
		} else {
			r -= b
			b *= 2
		}
	}
	for i := l; i <= r; i++ {
		fmt.Fprint(out, i, " ")
	}
	fmt.Fprintln(out)
}

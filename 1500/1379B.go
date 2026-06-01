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
	var l, r, m int
	fmt.Fscan(in, &l, &r, &m)
	for x := l; x <= r; x++ {
		if x > m {
			t := x - m
			if t <= r-l {
				fmt.Fprintln(out, x, l, l+t)
				return
			}
		} else if m%x <= r-l {
			fmt.Fprintln(out, x, l+m%x, l)
			return
		} else if x-m%x <= r-l {
			fmt.Fprintln(out, x, l, l+x-m%x)
			return
		}
	}
}

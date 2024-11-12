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
	var s string
	fmt.Fscan(in, &s)
	if n&1 == 1 {
		fmt.Fprintln(out, -1)
		return
	}
	cur, f := 0, 0
	b := true
	for i := range s {
		if s[i] == '(' {
			cur++
		} else {
			cur--
		}
		if f < 0 && cur > 0 || f > 0 && cur < 0 {
			b = false
		} else if f == 0 {
			f = cur
		}
	}
	if cur != 0 {
		fmt.Fprintln(out, -1)
		return
	}
	if !b {
		cur, color := 0, 1
		fmt.Fprintln(out, 2)
		for j := 0; j < n; j++ {
			if s[j] == '(' {
				cur++
			} else {
				cur--
			}
			if cur == 1 {
				color = 1
			} else if cur == -1 {
				color = 2
			}
			fmt.Fprint(out, color, " ")

		}
		fmt.Fprintln(out)
		return
	}
	fmt.Fprintln(out, 1)
	for i := 0; i < n; i++ {
		fmt.Fprint(out, 1, " ")
	}
	fmt.Fprintln(out)
}

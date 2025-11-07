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
	a := make([]int, n)
	pre := make([]int, n+1)
	xor := 0
	for i := range a {
		fmt.Fscan(in, &a[i])
		pre[i+1] = pre[i] ^ a[i]
	}

	var s string
	fmt.Fscan(in, &s)
	for i := range s {
		if s[i] == '1' {
			xor ^= a[i]
		}
	}

	var q, op, l, r int
	for fmt.Fscan(in, &q); q > 0; q-- {
		fmt.Fscan(in, &op, &l)
		if op == 1 {
			fmt.Fscan(in, &r)
			l--
			r--
			xor ^= pre[r+1] ^ pre[l]
		} else {
			if l == 1 {
				fmt.Fprint(out, xor, " ")
			} else {
				fmt.Fprint(out, xor^pre[n], " ")
			}
		}
	}

	fmt.Fprintln(out)
}

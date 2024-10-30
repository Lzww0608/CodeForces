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

	cur := 0
	pre := byte('2')
	for x := 2; x <= n; x++ {
		if s[x-2] != pre {
			cur = 1
		} else {
			cur++
		}
		pre = s[x-2]
		fmt.Fprint(out, x-cur, " ")
	}
	fmt.Fprintln(out)
}

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
	var s string
	fmt.Fscan(in, &n)
	var c byte = 0
	for i := 0; i <= n*2; i++ {
		fmt.Fscan(in, &s)
		for j := range s {
			c ^= s[j]
		}
	}
	fmt.Fprintln(out, string(c))
}

package main

import (
	"bufio"
	"fmt"
	"io"
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

func solve(in io.Reader, out io.Writer) {
	var n, k int
	fmt.Fscan(in, &n, &k)
	for i := 0; ; i++ {
		m := ((n >> i) + 1) / 2
		if k <= m {
			fmt.Fprintln(out, (k*2-1)<<i)
			return
		}
		k -= m
	}
}

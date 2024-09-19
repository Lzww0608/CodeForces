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

	var t, n int
	for fmt.Fscan(in, &t); t > 0; t-- {
		fmt.Fscan(in, &n)
		a := make([]int, n)
		p := make([]int, n)
		mex := n
		for i := range a {
			fmt.Fscan(in, &a[i])
		}
		for i := n - 1; i >= 0; i-- {
			p[i] = mex - a[i]
			mex = min(mex, p[i])
		}
		for _, x := range p {
			fmt.Fprint(out, x, " ")
		}
		fmt.Fprintln(out)
	}
}

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
	for i := range a {
		fmt.Fscan(in, &a[i])
	}

	f := make([]int, n)
	fa := 0
	for i := 1; i < n; i++ {
		if i > 1 && a[i-1] > a[i] {
			fa++
		}
		f[i] = f[fa] + 1
	}
	fmt.Fprintln(out, f[n-1])
}

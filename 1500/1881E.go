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
	f := make([]int, n+1)
	for i := range a {
		f[i] = n
		fmt.Fscan(in, &a[i])
	}

	for i := n - 1; i >= 0; i-- {
		f[i] = 1 + f[i+1]
		if i+a[i]+1 <= n {
			f[i] = min(f[i], f[i+a[i]+1])
		}
	}

	fmt.Fprintln(out, f[0])
}

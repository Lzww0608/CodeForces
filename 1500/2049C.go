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
	var n, x, y int
	fmt.Fscan(in, &n, &x, &y)

	a := make([]int, n)
	if n&1 == 1 || (y-x)&1 == 0 {
		a[x-1] = 2
	}
	for i := 1; i < n; i++ {
		j := (x + i - 1) % n
		if i&1 == 1 {
			a[j] = 1
		} else {
			a[j] = 0
		}
	}

	for i := range a {
		fmt.Fprintf(out, "%d ", a[i])
	}
	fmt.Fprintln(out)
}

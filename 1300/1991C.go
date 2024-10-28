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
	for i := 1; i < n; i++ {
		if a[0]&1 != a[i]&1 {
			fmt.Fprintln(out, -1)
			return
		}
	}
	fmt.Fprintln(out, 30+(a[0]&1)^1)
	for i := 29; i >= 0; i-- {
		fmt.Fprint(out, 1<<i, " ")
	}
	if a[0]&1 == 0 {
		fmt.Fprint(out, 1)
	}
	fmt.Fprintln(out)
}

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
	if n&1 == 0 {
		for i := 0; i < n; i += 2 {
			if a[i] > a[i+1] {
				a[i], a[i+1] = a[i+1], a[i]
			}
			if i > 0 && a[i] < a[i-1] {
				fmt.Fprintln(out, "NO")
				return
			}
		}
		fmt.Fprintln(out, "YES")
		return
	} else {
		for i := 1; i < n; i += 2 {
			if a[i] > a[i+1] {
				a[i], a[i+1] = a[i+1], a[i]
			}
			if a[i] < a[i-1] {
				fmt.Fprintln(out, "NO")
				return
			}
		}
		fmt.Fprintln(out, "YES")
		return
	}
}

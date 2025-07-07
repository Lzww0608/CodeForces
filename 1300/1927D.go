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
	a := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i])
	}
	left := make([]int, n+1)
	for i := 1; i <= n; i++ {
		if a[i] == a[i-1] {
			left[i] = left[i-1]
		} else {
			left[i] = i
		}
	}

	var q, l, r int
	for fmt.Fscan(in, &q); q > 0; q-- {
		fmt.Fscan(in, &l, &r)
		if left[l] == left[r] {
			fmt.Fprintln(out, -1, -1)
		} else {
			if a[r] != a[l] {
				fmt.Fprintln(out, l, r)
			} else {
				fmt.Fprintln(out, l, left[r]-1)
			}
		}
	}
	fmt.Fprintln(out)
}

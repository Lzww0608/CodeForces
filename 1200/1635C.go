package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
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

	if a[n-2] > a[n-1] {
		fmt.Fprintln(out, -1)
		return
	}

	if a[n-1] >= 0 {
		fmt.Fprintln(out, n-2)
		for i := 0; i < n-2; i++ {
			fmt.Fprintln(out, i+1, n-1, n)
		}
	} else if slices.IsSorted(a) {
		fmt.Fprintln(out, 0)
	} else {
		fmt.Fprintln(out, -1)
	}
}

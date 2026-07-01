package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
	var n, h int
	fmt.Fscan(in, &n, &h)
	a := make([]int, n)
	for i := range a {
		fmt.Fscan(in, &a[i])
	}
	sort.Ints(a)

	solve := func(b []int) int {
		sum := h
		for i, x := range a {
			if sum > x {
				sum += x / 2
			} else {
				for len(b) > 0 && sum <= x {
					sum *= b[0]
					b = b[1:]
				}

				if sum <= x {
					return i
				}
				sum += x / 2
			}
		}

		return len(a)
	}

	fmt.Fprintln(out, max(solve([]int{2, 2, 3}), solve([]int{2, 3, 2}), solve([]int{3, 2, 2})))
}

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, k int
	fmt.Fscan(in, &n, &k)
	a := make([]int, n)
	b := make([]int, k)
	for i := range a {
		fmt.Fscan(in, &a[i])
	}
	for i := range b {
		fmt.Fscan(in, &b[i])
	}
	if k > 1 {
		fmt.Fprintln(out, "Yes")
		return
	}
	pre := -1
	l, r := math.MinInt, math.MaxInt
	for i, x := range a {
		if x != 0 {
			if x < pre {
				fmt.Fprintln(out, "Yes")
				return
			}
			pre = x
		} else {
			if i > 0 {
				l = a[i-1]
			}
			if i < n-1 {
				r = a[i+1]
			}
		}
	}
	if b[0] > l && b[0] < r {
		fmt.Fprintln(out, "No")
		return
	}
	fmt.Fprintln(out, "Yes")
}

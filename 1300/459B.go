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

	var n int
	fmt.Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		fmt.Fscan(in, &a[i])
	}
	sort.Ints(a)
	x, y := 1, 1
	if a[0] == a[n-1] {
		fmt.Fprintln(out, 0, n*(n-1)/2)
		return
	}

	for i := 1; i < n; i++ {
		if a[i] == a[0] {
			x++
		}
	}
	for i := n - 2; i >= 0; i-- {
		if a[i] == a[n-1] {
			y++
		}
	}
	fmt.Fprintln(out, a[n-1]-a[0], x*y)
}

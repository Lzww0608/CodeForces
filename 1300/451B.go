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

	var n int
	fmt.Fscan(in, &n)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
	}

	i := 1
	for i < n && a[i] > a[i-1] {
		i++
	}
	if i == n {
		fmt.Fprintln(out, "yes")
		fmt.Fprintln(out, 1, 1)
		return
	}

	j := i
	for j < n && a[j] < a[j-1] {
		j++
	}
	for p, q := i-1, j-1; p < q; p, q = p+1, q-1 {
		a[p], a[q] = a[q], a[p]
	}

	for i := 1; i < n; i++ {
		if a[i] < a[i-1] {
			fmt.Fprintln(out, "no")
			return
		}
	}
	fmt.Fprintln(out, "yes")
	fmt.Fprintln(out, i, j)
}

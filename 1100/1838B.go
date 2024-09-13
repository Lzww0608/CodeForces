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

	var t, n int
	for fmt.Fscan(in, &t); t > 0; t-- {
		fmt.Fscan(in, &n)
		a := make([]int, n)
		for i := range a {
			fmt.Fscan(in, &a[i])
		}

		x, y, z := slices.Index(a, 1)+1, slices.Index(a, 2)+1, slices.Index(a, n)+1
		if z > max(x, y) {
			fmt.Fprintln(out, z, max(x, y))
		} else if z < min(x, y) {
			fmt.Fprintln(out, z, min(x, y))
		} else {
			fmt.Fprintln(out, 1, 1)
		}
	}
}

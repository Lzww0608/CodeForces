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

	var t, n int

	for fmt.Fscan(in, &t); t > 0; t-- {
		fmt.Fscan(in, &n)
		a := make([]int, n)
		for i := range a {
			fmt.Fscan(in, &a[i])
		}
		sort.Ints(a)
		i := n / 2
		for k := 1; i >= 0 && i < n; k++ {
			fmt.Fprint(out, a[i], " ")
			if k&1 == 1 {
				i -= k
			} else {
				i += k
			}
		}
		fmt.Fprintln(out)
	}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}

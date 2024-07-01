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
		if sort.IntsAreSorted(a) {
			fmt.Fprintln(out, 0)
		} else if a[0] == 1 || a[n-1] == n {
			fmt.Fprintln(out, 1)
		} else if a[0] == n && a[n-1] == 1 {
			fmt.Fprintln(out, 3)
		} else {
			fmt.Fprintln(out, 2)
		}

	}
}

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
		sort.Slice(a, func(i, j int) bool {
			return a[i] > a[j]
		})

		for _, x := range a {
			fmt.Fprint(out, x, " ")
		}
		fmt.Fprintln(out)
	}

}

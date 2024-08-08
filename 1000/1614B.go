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
		ans := make([]int, n)
		a := make([]int, n)
		pos := make([]int, n)
		for i := range a {
			fmt.Fscan(in, &a[i])
			pos[i] = i
		}

		sort.Slice(pos, func(i, j int) bool {
			return a[pos[i]] > a[pos[j]]
		})

		k, sum := 1, 0
		for _, i := range pos {
			ans[i] = k
			sum += max(k, -k) * 2 * a[i]
			if k > 0 {
				k = -k
			} else {
				k = -k + 1
			}
		}
		fmt.Fprintln(out, sum)
		fmt.Fprint(out, 0, " ")
		for _, x := range ans {
			fmt.Fprint(out, x, " ")
		}
		fmt.Fprintln(out)
	}
}

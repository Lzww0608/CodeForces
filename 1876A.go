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
		a := make([][2]int, n)
		for i := range a {
			fmt.Fscan(in, &a[i][0])
			a[i][1] = i
		}
		b := make([]int, n)
		sort.Slice(a, func(i, j int) bool {
			return a[i][0] < a[j][0]
		})
		for i := range a {
			b[a[i][1]] = n
			n--
		}
		for _, x := range b {
			fmt.Fprint(out, x, " ")
		}
		fmt.Fprintln(out)

	}
}

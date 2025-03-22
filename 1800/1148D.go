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
	a := make([][3]int, 0, n)
	b := make([][3]int, 0, n)
	for i := 0; i < n; i++ {
		var x, y int
		fmt.Fscan(in, &x, &y)
		if x < y {
			a = append(a, [3]int{x, y, i})
		} else {
			b = append(b, [3]int{x, y, i})
		}
	}

	sort.Slice(a, func(i, j int) bool {
		return a[i][0] > a[j][0]
	})
	sort.Slice(b, func(i, j int) bool {
		return b[i][1] < b[j][1]
	})

	if len(a) > len(b) {
		fmt.Fprintln(out, len(a))
		for _, v := range a {
			fmt.Fprint(out, v[2]+1, " ")
		}
	} else {
		fmt.Fprintln(out, len(b))
		for _, v := range b {
			fmt.Fprint(out, v[2]+1, " ")
		}
	}
}

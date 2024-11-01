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

	var t int
	for fmt.Fscan(in, &t); t > 0; t-- {
		solve(in, out)
	}
}

func solve(in *bufio.Reader, out *bufio.Writer) {
	var n int
	fmt.Fscan(in, &n)
	a := make([][2]int, n)
	for i := range a {
		fmt.Fscan(in, &a[i][0], &a[i][1])
	}
	sort.Slice(a, func(i, j int) bool {
		return a[i][0]+a[i][1] < a[j][0]+a[j][1]
	})
	for _, v := range a {
		fmt.Fprint(out, v[0], v[1], " ")
	}
	fmt.Fprintln(out)
}

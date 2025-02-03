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
	a := make([]int, n)
	for i := range a {
		fmt.Fscan(in, &a[i])
		a[i] += i + 1
	}
	sort.Slice(a, func(i, j int) bool {
		return a[i] > a[j]
	})
	for i := 1; i < n; i++ {
		a[i] = min(a[i], a[i-1]-1)
	}
	for _, x := range a {
		fmt.Fprint(out, x, " ")
	}
	fmt.Fprintln(out)
}

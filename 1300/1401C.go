package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
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
	b := make([]int, n)
	for i := range a {
		fmt.Fscan(in, &a[i])
		b[i] = a[i]
	}
	sort.Ints(b)
	x := slices.Min(a)
	if x == 1 {
		fmt.Fprintln(out, "YES")
		return
	}
	for i := 0; i < n; i++ {
		if a[i]%x == 0 {
			a[i] = 0
		}
		if b[i]%x == 0 {
			b[i] = 0
		}
		if a[i] != b[i] {
			fmt.Fprintln(out, "NO")
			return
		}
	}
	fmt.Fprintln(out, "YES")
}

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
	p := make([]int, n)
	for i := range a {
		fmt.Fscan(in, &a[i])
		a[i]--
		p[i] = i
	}

	sort.Slice(p, func(i, j int) bool {
		return a[p[i]] > a[p[j]]
	})

	for _, x := range p {
		fmt.Fprint(out, x+1, " ")
	}
	fmt.Fprintln(out)
}

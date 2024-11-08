package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)
	a := make([]int, n)
	b := make([]int, n)
	pos := make([]int, n)
	for i := range a {
		fmt.Fscan(in, &a[i])
		a[i]--
		pos[a[i]] = i
	}
	cnt := make([]int, n)
	for i := range b {
		fmt.Fscan(in, &b[i])
		b[i]--
		cnt[(i-pos[b[i]]+n)%n]++
	}
	fmt.Fprintln(out, slices.Max(cnt))
}

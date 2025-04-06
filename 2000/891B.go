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
	a := make([]int, n)
	id := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
		id[i] = i
	}
	sort.Slice(id, func(i, j int) bool {
		return a[id[i]] < a[id[j]]
	})

	b := make([]int, n)
	for i := 0; i < n; i++ {
		b[id[i%n]] = a[id[(i+1)%n]]
	}

	for _, x := range b {
		fmt.Fprint(out, x, " ")
	}
}

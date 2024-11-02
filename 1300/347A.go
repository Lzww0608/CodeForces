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
	for i := range a {
		fmt.Fscan(in, &a[i])
	}
	sort.Ints(a)
	fmt.Fprint(out, a[n-1], " ")
	for _, x := range a[1 : n-1] {
		fmt.Fprint(out, x, " ")
	}
	fmt.Fprint(out, a[0])
}

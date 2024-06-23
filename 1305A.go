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
		a := make([]int, n)
		b := make([]int, n)
		for i := range a {
			fmt.Fscan(in, &a[i])
		}
		for i := range b {
			fmt.Fscan(in, &b[i])
		}
		sort.Ints(a)
		sort.Ints(b)
		for i := range a {
			fmt.Fprint(out, a[i], " ")
		}
		fmt.Fprintln(out)
		for i := range b {
			fmt.Fprint(out, b[i], " ")
		}
		fmt.Fprintln(out)
	}
}

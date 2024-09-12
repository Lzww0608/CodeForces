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
next:
	for fmt.Fscan(in, &t); t > 0; t-- {
		fmt.Fscan(in, &n)
		a := make([]int, n)
		for i := range a {
			fmt.Fscan(in, &a[i])
		}
		if n&1 == 1 {
			fmt.Fprintln(out, "NO")
			continue
		}
		sort.Ints(a)
		for i := 0; i < n/2; i++ {
			if a[i] >= a[i+n/2] || (i > 0 && a[i] >= a[i+n/2-1]) {
				fmt.Fprintln(out, "NO")
				continue next
			}
		}
		fmt.Fprintln(out, "YES")
		for i := 0; i < n/2; i++ {
			fmt.Fprint(out, a[i], a[i+n/2], " ")
		}
		fmt.Fprintln(out)
	}
}

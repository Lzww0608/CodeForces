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

	var t, n, x int
	for fmt.Fscan(in, &t); t > 0; t-- {
		fmt.Fscan(in, &n, &x)
		a := make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Fscan(in, &a[i])
		}
		if n >= x*2 {
			fmt.Fprintln(out, "YES")
			continue
		} else if ok := sort.SliceIsSorted(a[n-x+1:x], func(i, j int) bool {
			return a[i] < a[j]
		}); ok {
			fmt.Fprintln(out, "YES")
			continue
		}
		fmt.Fprintln(out, "NO")
	}
}

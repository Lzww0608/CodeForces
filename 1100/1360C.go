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
		odd, even := 0, 0
		for i := range a {
			fmt.Fscan(in, &a[i])
			if a[i]&1 == 1 {
				odd++
			} else {
				even++
			}
		}
		if odd&1 == 0 {
			fmt.Fprintln(out, "YES")
			continue
		}
		sort.Ints(a)
		for i := 1; i < n; i++ {
			if a[i] == a[i-1]+1 {
				fmt.Fprintln(out, "YES")
				continue next
			}
		}
		fmt.Fprintln(out, "NO")
	}
}

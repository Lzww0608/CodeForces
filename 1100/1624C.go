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
		sort.Ints(a)
		cnt := make([]int, n+1)
		for _, x := range a {
			for x > n || cnt[x] != 0 {
				x >>= 1
			}
			if x == 0 || cnt[x] != 0 {
				fmt.Fprintln(out, "NO")
				continue next
			}
			cnt[x]++
		}
		fmt.Fprintln(out, "YES")
	}
}

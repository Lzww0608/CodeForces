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
		if a[0] == a[1] {
			if a[0] == 1 {
				fmt.Fprintln(out, "NO")
			} else {
				for i := 2; i < n; i++ {
					if a[i]%a[0] < a[0] && a[i]%a[0] != 0 {
						fmt.Fprintln(out, "YES")
						continue next
					}
				}
				fmt.Fprintln(out, "NO")
			}
		} else {
			fmt.Fprintln(out, "YES")
		}
	}
}

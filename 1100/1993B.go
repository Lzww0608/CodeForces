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
		cnt, mx := 0, 0

		for i := range a {
			fmt.Fscan(in, &a[i])
			if a[i]&1 == 0 {
				cnt++
			} else {
				mx = max(mx, a[i])
			}

		}
		if cnt == n || cnt == 0 {
			fmt.Fprintln(out, 0)
			continue
		}

		ans := cnt
		sort.Ints(a)
		for _, x := range a {
			if x&1 == 0 {
				if x < mx {
					mx += x
				} else {
					ans++
					break
				}
			}
		}
		fmt.Fprintln(out, ans)
	}
}

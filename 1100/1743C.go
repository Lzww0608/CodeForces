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

	var t, n int
	for fmt.Fscan(in, &t); t > 0; t-- {
		fmt.Fscan(in, &n)
		var s string
		a := make([]int, n)
		fmt.Fscan(in, &s)
		for i := range a {
			fmt.Fscan(in, &a[i])
		}
		ans := 0
		l, r := 0, 0
		for r < n {
			l = r
			for r < n && s[r] == '1' {
				ans += a[r]
				r++
			}
			if r > l && r <= n && l > 0 {
				mn := slices.Min(a[l:r])
				if mn < a[l-1] {
					ans += a[l-1] - mn
				}
			}
			r++
		}
		fmt.Fprintln(out, ans)
	}
}

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var t, n int
	for fmt.Fscan(in, &t); t > 0; t-- {
		fmt.Fscan(in, &n)
		a := make([]int, n)
		for i := range a {
			fmt.Fscan(in, &a[i])
		}
		ans := n
		if n >= 4 && n&1 == 0 {
			x, y := a[0], a[1]
			f := true
			for i := 2; i < n; i += 2 {
				if i+1 < n && x == a[i] && y == a[i+1] {
					continue
				} else {
					f = false
					break
				}
			}
			if f {
				ans -= (n - 2) / 2
			}
		}
		fmt.Fprintln(out, ans)
	}
}

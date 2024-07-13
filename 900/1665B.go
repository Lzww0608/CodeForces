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

	var t, n, x int
	for fmt.Fscan(in, &t); t > 0; t-- {
		fmt.Fscan(in, &n)
		m := map[int]int{}
		mx := 0
		for i := 0; i < n; i++ {
			fmt.Fscan(in, &x)
			m[x]++
			mx = max(mx, m[x])
		}
		ans := 0
		for mx < n {
			ans++
			if mx*2 <= n {
				ans += mx
			} else {
				ans += n - mx
			}
			mx = mx * 2
		}
		fmt.Fprintln(out, ans)
	}
}

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var t, n, c, d, x int
next:
	for fmt.Fscan(in, &t); t > 0; t-- {
		fmt.Fscan(in, &n, &c, &d)
		m := map[int]int{}
		mn := math.MaxInt32
		for i := 0; i < n*n; i++ {
			fmt.Fscan(in, &x)
			m[x]++
			mn = min(mn, x)
		}

		for i := 0; i < n; i++ {
			x = mn + i*c
			for j := 0; j < n; j++ {
				if m[x+j*d] <= 0 {
					fmt.Fprintln(out, "NO")
					continue next
				}
				m[x+j*d]--
			}
		}
		fmt.Fprintln(out, "YES")
	}
}

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

	var t, n, x int
	for fmt.Fscan(in, &t); t > 0; t-- {
		fmt.Fscan(in, &n)
		cnt, mx := 0, math.MinInt
		for i := 0; i < n; i++ {
			fmt.Fscan(in, &x)
			if x > 0 {
				mx = max(mx, x)
				cnt++
			}
		}
		if cnt == 0 {
			fmt.Fprintln(out, 1)
		} else if cnt >= n/2 {
			fmt.Fprintln(out, 0)
		} else if mx == 1 {
			fmt.Fprintln(out, 2)
		} else {
			fmt.Fprintln(out, 1)
		}
	}
}

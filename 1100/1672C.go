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

	var t, n int
	for fmt.Fscan(in, &t); t > 0; t-- {
		fmt.Fscan(in, &n)
		a := make([]int, n)
		for i := range a {
			fmt.Fscan(in, &a[i])
		}
		mn, mx := math.MaxInt32, math.MinInt32
		for i := 0; i < n-1; i++ {
			if a[i] == a[i+1] {
				mn = min(mn, i)
				mx = max(mx, i)
			}
		}
		if mn >= mx {
			fmt.Fprintln(out, 0)
		} else {
			fmt.Fprintln(out, max(1, mx-mn-1))
		}

	}
}

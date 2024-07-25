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

	var t, m, n, x int
	for fmt.Fscan(in, &t); t > 0; t-- {
		fmt.Fscan(in, &m)
		a := make([][2]int, m)
		for i := range a {
			fmt.Fscan(in, &n)
			a[i][0], a[i][1] = math.MaxInt, math.MaxInt
			for j := 0; j < n; j++ {
				fmt.Fscan(in, &x)
				if x < a[i][0] {
					a[i][0], a[i][1] = x, a[i][0]
				} else if x < a[i][1] {
					a[i][1] = x
				}
			}
		}
		ans, mn, t := 0, math.MaxInt, math.MaxInt
		for _, v := range a {
			ans += v[1]
			mn = min(mn, v[0])
			t = min(t, v[1])
		}
		fmt.Fprintln(out, ans+mn-t)
	}
}

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

	var n, l, r, x int
	fmt.Fscan(in, &n, &l, &r, &x)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
	}

	ans := 0
	for s := 1; s < (1 << n); s++ {
		sum := 0
		mn, mx := math.MaxInt32, math.MinInt32
		for i := 0; i < n; i++ {
			if s&(1<<i) != 0 {
				sum += a[i]
				mn = min(mn, a[i])
				mx = max(mx, a[i])
			}
		}
		if sum >= l && sum <= r && mx-mn >= x {
			ans++
		}
	}

	fmt.Fprintln(out, ans)
}

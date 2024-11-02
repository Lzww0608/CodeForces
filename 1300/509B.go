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

	var n, k int
	fmt.Fscan(in, &n, &k)
	a := make([]int, n)
	mn, mx := math.MaxInt32, math.MinInt32
	for i := range a {
		fmt.Fscan(in, &a[i])
		mn = min(mn, a[i])
		mx = max(mx, a[i])
	}
	if mx-mn > k {
		fmt.Fprintln(out, "NO")
		return
	}
	fmt.Fprintln(out, "YES")
	for _, x := range a {
		for t := 0; x > 0; t = (t + 1) % k {
			fmt.Fprint(out, t+1, " ")
			x--
		}
		fmt.Fprintln(out)
	}
}

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

	var n int
	fmt.Fscan(in, &n)
	a := make([]int, n)
	mn := math.MaxInt32
	for i := range a {
		fmt.Fscan(in, &a[i])
		mn = min(mn, a[i])
	}
	ans := 0
	for i := 0; i < n; i++ {
		ans += mn
		a[i] -= mn
	}
	mx, cur := 0, 0
	for i := 0; i < n*2; i++ {
		if a[i%n] == 0 {
			mx = max(mx, cur)
			cur = 0
		} else {
			cur++
		}
	}
	fmt.Fprintln(out, mx+ans)
}

package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var t int
	for fmt.Fscan(in, &t); t > 0; t-- {
		solve(in, out)
	}
}

func solve(in *bufio.Reader, out *bufio.Writer) {
	var n, x, y int
	fmt.Fscan(in, &n)
	a := make([]int, n)
	Or := 0
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &x)
		for j := 0; j < x; j++ {
			fmt.Fscan(in, &y)
			a[i] |= 1 << y
		}
		Or |= a[i]
	}

	ans := 0
	for i := 0; i <= 50; i++ {
		if (Or>>i)&1 == 1 {
			v := 0
			for j := range a {
				if a[j]&(1<<i) == 0 {
					v |= a[j]
				}
			}
			ans = max(ans, bits.OnesCount(uint(v)))
		}
	}
	fmt.Fprintln(out, ans)
}

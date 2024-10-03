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
		cnt, mn := 0, math.MaxInt32
		for i := 0; i < n; i++ {
			fmt.Fscan(in, &x)
			if x&1 == 0 {
				cnt++
				tmp := 0
				for x&1 == 0 {
					x >>= 1
					tmp++
				}
				mn = min(mn, tmp)
			}
		}

		if cnt < n {
			fmt.Fprintln(out, cnt)
		} else {
			fmt.Fprintln(out, cnt-1+mn)
		}

	}
}

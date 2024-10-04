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

	var t, r int

	for fmt.Fscan(in, &t); t > 0; t-- {
		fmt.Fscan(in, &r)
		ans := make([]int, r)
		var dfs func(r int)
		dfs = func(r int) {
			if r < 0 {
				return
			}
			s := int(math.Sqrt(float64(r * 2)))
			s *= s
			l := s - r
			dfs(l - 1)
			for l <= r {
				ans[l], ans[r] = r, l
				l++
				r--
			}
		}
		dfs(r - 1)
		for _, x := range ans {
			fmt.Fprint(out, x, " ")
		}
		fmt.Fprintln(out)
	}
}

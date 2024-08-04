package main

import (
	"bufio"
	"fmt"
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
		b := make([]int, n)
		for i := range a {
			fmt.Fscan(in, &a[i])
		}
		for i := range b {
			fmt.Fscan(in, &b[i])
		}
		ans := 1
		ma := map[int]int{}
		mb := map[int]int{}
		cur, x := 0, -1
		for _, t := range a {
			if x == t {
				cur++
			} else {
				x, cur = t, 1
			}
			ma[x] = max(ma[x], cur)
		}
		cur, x = 0, -1
		for _, t := range b {
			if x == t {
				cur++
			} else {
				x, cur = t, 1
			}
			mb[x] = max(mb[x], cur)
			ans = max(ans, mb[x]+ma[x])
		}

		for k, v := range ma {
			ans = max(ans, v+mb[k])
		}
		fmt.Fprintln(out, ans)
	}

}

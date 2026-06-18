package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)
	c := make([]pair, n)
	for i := range c {
		fmt.Fscan(in, &c[i].a, &c[i].b)
	}

	sort.Slice(c, func(i, j int) bool {
		return c[i].b < c[j].b
	})

	ans := 0
	for l, r, x := 0, n-1, 0; l <= r; {
		if c[l].b > x {
			t := min(c[l].b-x, c[r].a)
			ans += t * 2
			x += t
			if c[r].a -= t; c[r].a == 0 {
				r--
			}
		} else {
			ans += c[l].a
			x += c[l].a
			l++
		}
	}

	fmt.Fprintln(out, ans)
}

type pair struct {
	a, b int
}

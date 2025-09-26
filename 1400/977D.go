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
	type pair struct {
		x, y int
	}

	calc := func(x int) (cnt int) {
		for x%3 == 0 {
			cnt++
			x /= 3
		}

		return
	}

	a := make([]pair, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i].y)
		a[i].x = -calc(a[i].y)
	}

	sort.Slice(a, func(i, j int) bool {
		if a[i].x == a[j].x {
			return a[i].y < a[j].y
		}

		return a[i].x < a[j].x
	})

	for _, v := range a {
		fmt.Fprint(out, v.y, " ")
	}
}

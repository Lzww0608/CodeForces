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

	var n, S, T int
	fmt.Fscan(in, &n, &S)
	a := make([]int, n)
	b := make([]int, n)
	for i := range a {
		fmt.Fscan(in, &a[i])
	}

	check := func(k int) (ok bool, t int) {
		for i := range n {
			b[i] = a[i] + k*(i+1)
		}
		sort.Ints(b)
		for _, x := range b[:k] {
			t += x
			if t > S {
				return false, 0
			}
		}
		return true, t
	}

	l, r := -1, n+1
	for l+1 < r {
		mid := l + ((r - l) >> 1)
		if ok, t := check(mid); ok {
			l, T = mid, t
		} else {
			r = mid
		}
	}

	fmt.Fprintln(out, l, T)
}

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

	var n, m int
	fmt.Fscan(in, &n, &m)
	a := make([]int, n)
	sum := 0
	for i := range a {
		fmt.Fscan(in, &a[i])
		sum += a[i]
	}

	sort.Ints(a)
	l, r := 0, n-1
	w := n
	for l < r {
		if a[l]+a[r] <= m {
			if l+1 == r {
				w -= 1
				break
			}
			l++
			r--
			w -= 2
		} else {
			r--
		}
	}
	fmt.Fprintln(out, sum+w)
}

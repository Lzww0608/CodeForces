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

	var n, q int
	fmt.Fscan(in, &n, &q)
	a := make([]int, n)
	for i := range a {
		fmt.Fscan(in, &a[i])
	}

	f := make([]int, n)
	for i := 1; i < n-1; i++ {
		if a[i-1] >= a[i] && a[i] >= a[i+1] {
			f[i] = f[i-1] + 1
		} else {
			f[i] = f[i-1]
		}
	}

	var l, r int
	for range q {
		fmt.Fscan(in, &l, &r)
		if l == r {
			fmt.Fprintln(out, 1)
			continue
		}
		l--
		r--
		fmt.Fprintln(out, r-l+1-f[r-1]+f[l])
	}
}

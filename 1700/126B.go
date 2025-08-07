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

	var s string
	fmt.Fscan(in, &s)
	n := len(s)
	z := make([]int, n)
	for i, l, r := 0, 0, 0; i < n; i++ {
		z[i] = max(min(z[i-l], r-i+1), 0)
		for i+z[i] < n && s[i+z[i]] == s[z[i]] {
			l, r = i, i+z[i]
			z[i]++
		}
	}

	a := []int{}
	for l := 1; l < n; l++ {
		if z[n-l] >= l {
			a = append(a, l)
		}
	}

	for i := len(a) - 1; i >= 0; i-- {
		l := a[i]
		for j := 1; j < n-l; j++ {
			if z[j] >= l {
				fmt.Fprintln(out, s[:l])
				return
			}
		}
	}

	fmt.Fprintln(out, "Just a legend")
}

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

	var t, n, k int
	for fmt.Fscan(in, &t); t > 0; t-- {
		fmt.Fscan(in, &n, &k)
		a := make([]int, n)
		b := make([]int, n)
		for i := range a {
			fmt.Fscan(in, &a[i])
		}
		for i := range b {
			fmt.Fscan(in, &b[i])
		}
		sort.Ints(a)
		sort.Ints(b)
		l := []int{}
		r := []int{}
		for i, j := 0, 0; i < n && j < k; i++ {
			if i < n-1 && a[i] == a[i+1] {
				l = append(l, a[i:i+2]...)
				j++
			}
		}
		for i, j := 0, 0; i < n && j < k; i++ {
			if i < n-1 && b[i] == b[i+1] {
				r = append(r, b[i:i+2]...)
				j++
			}
		}

		for i, j := 0, 0; i < n && j < n && len(l) < k*2; {
			if a[i] > b[j] {
				j++
			} else if a[i] < b[j] {
				i++
			} else {
				l = append(l, a[i])
				r = append(r, b[j])
				i++
				j++
			}
		}

		for _, x := range l {
			fmt.Fprint(out, x, " ")
		}
		fmt.Fprintln(out)
		for _, x := range r {
			fmt.Fprint(out, x, " ")
		}
		fmt.Fprintln(out)

	}
}

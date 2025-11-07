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

	var t int
	for fmt.Fscan(in, &t); t > 0; t-- {
		solve(in, out)
	}
}

func solve(in *bufio.Reader, out *bufio.Writer) {
	var n int
	fmt.Fscan(in, &n)
	d := make([]int, n)
	for i := range d {
		fmt.Fscan(in, &d[i])
	}

	sort.Ints(d)
	x := d[0] * d[n-1]

	check := func() bool {
		a := []int{}
		for i := 2; i*i <= x; i++ {
			if x%i == 0 {
				a = append(a, i)
				if i*i != x {
					a = append(a, x/i)
				}
			}
		}
		if len(a) != n {
			return false
		}
		sort.Ints(a)
		for i := range a {
			if a[i] != d[i] {
				return false
			}
		}

		return true
	}

	if check() {
		fmt.Fprintln(out, x)
	} else {
		fmt.Fprintln(out, -1)
	}
}

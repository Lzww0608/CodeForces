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
	a := make([]int, n)
	for i := range a {
		fmt.Fscan(in, &a[i])
	}
	sort.Ints(a)
	cnt := 0
	for i := range a {
		if i == 0 || a[i] > a[i-1] {
			cnt++
		}
	}
	for i := 1; i < n; i++ {
		if a[i] == a[i-1] {
			cnt++
			for i < n && a[i] == a[i-1] {
				i++
			}
		}
	}

	fmt.Fprintln(out, (cnt+1)/2)
}

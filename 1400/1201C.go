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

	var n, k int
	fmt.Fscan(in, &n, &k)
	a := make([]int, n)
	for i := range a {
		fmt.Fscan(in, &a[i])
	}
	sort.Ints(a)
	a = a[n/2:]
	n, i := len(a), 1
	for i < n {
		if a[i] != a[i-1] {
			if k < i*(a[i]-a[i-1]) {
				break
			}
			k -= i * (a[i] - a[i-1])
		}
		i++
	}

	fmt.Fprintln(out, a[i-1]+k/i)
}

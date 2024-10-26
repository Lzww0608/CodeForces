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
	var n, w int
	fmt.Fscan(in, &n, &w)
	a := make([]int, n)
	id := make([]int, n)
	sum := 0
	for i := range a {
		fmt.Fscan(in, &a[i])
		id[i] = i
		sum += a[i]
	}
	sort.Slice(id, func(i, j int) bool {
		return a[id[i]] > a[id[j]]
	})
	half := (w + 1) / 2
	if a[id[n-1]] > w || sum < half {
		fmt.Fprintln(out, -1)
		return
	}

	sum = 0
	ans := []int{}
	for _, i := range id {
		if sum+a[i] <= w {
			sum += a[i]
			ans = append(ans, i+1)
			if sum >= half {
				break
			}
		}
	}
	if sum < half {
		fmt.Fprintln(out, -1)
		return
	}
	fmt.Fprintln(out, len(ans))
	for _, x := range ans {
		fmt.Fprint(out, x, " ")
	}
	fmt.Fprintln(out)
}

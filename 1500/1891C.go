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
	ans, sum := 0, 0
	for i := range a {
		fmt.Fscan(in, &a[i])
		sum += a[i]
	}
	sort.Slice(a, func(i, j int) bool {
		return a[i] > a[j]
	})
	sum /= 2
	for i := 0; i < n; i++ {
		if sum >= a[i] {
			sum -= a[i]
			a[i] = 0
		} else if sum > 0 {
			a[i] -= sum
			sum = 0
		} else {
			break
		}
		ans++
	}
	for _, x := range a {
		ans += x
	}
	fmt.Fprintln(out, ans)
}

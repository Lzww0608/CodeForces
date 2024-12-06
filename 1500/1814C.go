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
	var n, s1, s2 int
	fmt.Fscan(in, &n, &s1, &s2)
	a := make([][2]int, n)
	for i := range a {
		fmt.Fscan(in, &a[i][0])
		a[i][1] = i + 1
	}
	sort.Slice(a, func(i, j int) bool {
		return a[i][0] > a[j][0]
	})

	ans1 := []int{}
	ans2 := []int{}
	sum1, sum2 := 0, 0
	for i := 0; i < n; i++ {
		if sum1+s1 < sum2+s2 {
			sum1 += s1
			ans1 = append(ans1, a[i][1])
		} else {
			sum2 += s2
			ans2 = append(ans2, a[i][1])
		}
	}
	fmt.Fprint(out, len(ans1), " ")
	for _, x := range ans1 {
		fmt.Fprint(out, x, " ")
	}
	fmt.Fprintln(out)
	fmt.Fprint(out, len(ans2), " ")
	for _, x := range ans2 {
		fmt.Fprint(out, x, " ")
	}
	fmt.Fprintln(out)
}

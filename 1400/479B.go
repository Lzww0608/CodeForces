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
	a := make([][2]int, n)
	for i := range a {
		fmt.Fscan(in, &a[i][0])
		a[i][1] = i + 1
	}
	ans := [][]int{}
	sort.Slice(a, func(i, j int) bool {
		return a[i][0] < a[j][0]
	})
	for i := 0; i < k && a[0][0]+1 < a[n-1][0]; i++ {
		ans = append(ans, []int{a[n-1][1], a[0][1]})
		a[n-1][0]--
		a[0][0]++
		sort.Slice(a, func(i, j int) bool {
			return a[i][0] < a[j][0]
		})
	}

	fmt.Fprintln(out, a[n-1][0]-a[0][0], len(ans))
	for _, v := range ans {
		fmt.Fprintln(out, v[0], v[1])
	}
}

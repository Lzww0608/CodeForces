package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

const MOD int = 998244353

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, m, k int
	fmt.Fscan(in, &n, &m, &k)
	a := make([][2]int, n)
	for i := range a {
		fmt.Fscan(in, &a[i][0])
		a[i][1] = i
	}

	sort.Slice(a, func(i, j int) bool {
		return a[i][0] > a[j][0] || a[i][0] == a[j][0] && a[i][1] < a[j][1]
	})
	sum := 0
	id := make([]int, m*k)
	for i := 0; i < m*k; i++ {
		sum += a[i][0]
		id[i] = a[i][1]
	}
	sort.Ints(id)

	div := make([]int, k-1)
	for i := range div {
		div[i] = id[(i+1)*m-1]
	}
	fmt.Fprintln(out, sum)
	for _, x := range div {
		fmt.Fprint(out, x+1, " ")
	}
}

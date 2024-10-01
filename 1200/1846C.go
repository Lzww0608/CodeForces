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
	var n, m, h int
	fmt.Fscan(in, &n, &m, &h)
	a := make([][]int, n)
	id := make([]int, n)
	for i := range a {
		a[i] = make([]int, m)
		for j := range a[i] {
			fmt.Fscan(in, &a[i][j])
		}
		sort.Ints(a[i])
	}
	ans := 0
	for i := 0; i < n; i++ {
		id[i] = i
		cnt, pen := 0, 0
		sum := 0
		for j := 0; j < m && sum+a[i][j] <= h; j++ {
			cnt++
			sum += a[i][j]
			pen += sum
		}
		a[i] = []int{cnt, pen}
		if cnt > a[0][0] || (cnt == a[0][0] && pen < a[0][1]) {
			ans++
		}
	}
	fmt.Fprintln(out, ans+1)
}

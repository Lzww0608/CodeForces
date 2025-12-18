package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)
	a := make([]int, n)
	b := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
	}
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &b[i])
	}

	w := bits.Len(uint(n))
	sta := make([][]int, w)
	stb := make([][]int, w)
	for i := range w {
		sta[i] = make([]int, n)
		stb[i] = make([]int, n)
	}

	sta[0] = a
	stb[0] = b
	for i := 1; i < w; i++ {
		for j := range n - (1 << i) + 1 {
			sta[i][j] = max(sta[i-1][j], sta[i-1][j+(1<<(i-1))])
			stb[i][j] = min(stb[i-1][j], stb[i-1][j+(1<<(i-1))])
		}
	}

	qa := func(l, r int) int {
		k := bits.Len(uint(r-l)) - 1
		return max(sta[k][l], sta[k][r-(1<<k)])
	}

	qb := func(l, r int) int {
		k := bits.Len(uint(r-l)) - 1
		return min(stb[k][l], stb[k][r-(1<<k)])
	}

	ans := 0
	for i := range n {
		l := sort.Search(n-i, func(j int) bool {
			return qa(i, i+j+1) >= qb(i, i+j+1)
		}) + i
		r := sort.Search(n-i, func(j int) bool {
			return qa(i, i+j+1) > qb(i, i+j+1)
		}) + i

		ans += r - l
	}

	fmt.Fprintln(out, ans)
}

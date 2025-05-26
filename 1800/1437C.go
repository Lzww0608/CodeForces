package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var q int
	for fmt.Fscan(in, &q); q > 0; q-- {
		solve(in, out)
	}

}

func solve(in *bufio.Reader, out *bufio.Writer) {
	var n int
	fmt.Fscan(in, &n)
	t := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &t[i])
	}

	f := make([][]int, n+1)
	for i := range f {
		f[i] = make([]int, n*2+1)
		for j := range f[i] {
			f[i][j] = math.MaxInt32
			f[0][j] = 0
		}
	}

	sort.Ints(t)
	for i := 1; i <= n; i++ {
		for j := 1; j <= n*2; j++ {
			f[i][j] = min(f[i][j-1], f[i-1][j-1]+abs(j-t[i-1]))
		}
	}

	ans := math.MaxInt32
	for _, x := range f[n] {
		ans = min(ans, x)
	}
	fmt.Fprintln(out, ans)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

const N int = 2001

var f [N][N]int

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		fmt.Fscan(in, &a[i])
	}

	sort.Ints(a)
	for i := range N {
		for j := range N {
			f[i][j] = -1
		}
	}

	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		if i > j {
			return 0
		}
		if f[i][j] != -1 {
			return f[i][j]
		}
		sum := a[j] - a[i]
		sum += min(dfs(i+1, j), dfs(i, j-1))
		f[i][j] = sum
		return sum
	}

	fmt.Fprintln(out, dfs(0, n-1))
}

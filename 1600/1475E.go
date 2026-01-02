package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

const N int = 1_001
const MOD int = 1_000_000_007

var comb [N][N]int

func init() {
	for i := range N {
		comb[i][0], comb[i][i] = 1, 1
		for j := i - 1; j > 0; j-- {
			comb[i][j] = comb[i-1][j] + comb[i-1][j-1]
			comb[i][j] %= MOD
		}
	}
}

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
	var n, k int
	fmt.Fscan(in, &n, &k)
	a := make([]int, n)
	for i := range a {
		fmt.Fscan(in, &a[i])
	}

	sort.Slice(a, func(i, j int) bool {
		return a[i] > a[j]
	})

	cnt := 0
	for i := 0; i < n && cnt < k; i++ {
		j := i
		for j < n && a[i] == a[j] {
			j++
		}
		cur := j - i
		if cnt+cur <= k {
			cnt += cur
		} else {
			d := k - cnt
			fmt.Fprintln(out, comb[cur][d])
			return
		}
		i = j - 1
	}

	fmt.Fprintln(out, 1)
}

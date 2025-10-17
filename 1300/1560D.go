package main

import (
	"bufio"
	"fmt"
	"os"
)

const MOD int = 1_000_000_007

var nums []string

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	for i := 1; i <= 1e18; i *= 2 {
		nums = append(nums, fmt.Sprintf("%d", i))
	}

	var t int
	for fmt.Fscan(in, &t); t > 0; t-- {
		solve(in, out)
	}
}

func lcs(t, s string) int {
	m, n := len(s), len(t)
	i, j := 0, 0

	for i < m && j < n {
		if s[i] == t[j] {
			i++
		}
		j++
	}

	return i
}

func calc(s, t string) int {
	m, n := len(s), len(t)
	return m + n - 2*lcs(s, t)
}

func solve(in *bufio.Reader, out *bufio.Writer) {
	var s string
	fmt.Fscan(in, &s)
	ans := len(s) + 1

	for _, t := range nums {
		ans = min(ans, calc(s, t))
	}

	fmt.Fprintln(out, ans)
}

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
	s := make([]string, n)
	id := make([]int, n)
	cnt := make([][5]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &s[i])
		id[i] = i
		for j := range s[i] {
			cnt[i][s[i][j]-'a']++
		}
		for j := range cnt[i] {
			cnt[i][j] = cnt[i][j]*2 - len(s[i])
		}
	}

	calc := func(t int) int {
		res := 0
		sort.Slice(id, func(i, j int) bool {
			return cnt[id[i]][t] > cnt[id[j]][t]
		})
		cur := 0
		for j, i := range id {
			cur += cnt[i][t]
			if cur > 0 {
				res = j + 1
			} else {
				break
			}
		}
		return res
	}

	ans := 0
	for i := range 5 {
		ans = max(ans, calc(i))
	}
	fmt.Fprintln(out, ans)
}

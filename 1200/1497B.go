package main

import (
	"bufio"
	"fmt"
	"os"
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
	var n, m, x int
	fmt.Fscan(in, &n, &m)
	cnt := make(map[int]int)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &x)
		cnt[x%m]++
	}
	ans := 0
	for k := 0; k+k <= m; k++ {
		v := cnt[k]
		t := (m - k) % m
		if cnt[t] == 0 {
			ans += v
		} else {
			if k != t && k != 0 {
				if v >= cnt[t] {
					ans += max(0, v-cnt[t]-1)
				} else {
					ans += cnt[t] - v - 1
				}
			}
			ans++
		}
		cnt[k] = 0
	}
	fmt.Fprintln(out, ans)
}

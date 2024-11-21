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
	var n int
	fmt.Fscan(in, &n)
	var s, t string
	fmt.Fscan(in, &s, &t)
	cnt := 0
	for i := 0; i < n; i++ {
		if s[i] != t[i] {
			cnt++
		}
	}
	ans := [][]int{}
	if cnt == n {
		ans = append(ans, []int{1, n})
		s = t
	} else if cnt != 0 {
		fmt.Fprintln(out, "NO")
		return
	}

	sum := make([]int, n+1)
	for i := 0; i < n; i++ {
		if s[i] == '1' {
			if i == 0 {
				ans = append(ans, []int{1, n})
				ans = append(ans, []int{2, n})
			} else {
				sum[i+1]++
				sum[i]++
			}
		}
	}

	for i, x := range sum {
		if x&1 == 1 {
			ans = append(ans, []int{1, i})
		}
	}

	fmt.Fprintln(out, "YES")
	fmt.Fprintln(out, len(ans))
	for _, v := range ans {
		fmt.Fprintln(out, v[0], v[1])
	}
}

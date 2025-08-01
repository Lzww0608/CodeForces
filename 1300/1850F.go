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
	a := make([]int, n)
	cnt := make([]int, n+1)
	for i := range a {
		fmt.Fscan(in, &a[i])
		if a[i] <= n {
			cnt[a[i]]++
		}
	}

	ans := 0
	for i := 1; i <= n; i++ {
		cur := 0
		for j := 1; j*j <= i; j++ {
			if i%j == 0 {
				cur += cnt[j]
				if j*j != i {
					cur += cnt[i/j]
				}
			}
		}
		ans = max(ans, cur)
	}

	fmt.Fprintln(out, ans)
}

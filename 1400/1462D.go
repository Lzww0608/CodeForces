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
	for i := range a {
		fmt.Fscan(in, &a[i])
	}

	ans := n
	pre := 0
	for i := range n {
		pre += a[i]
		cnt, cur := 1, 0
		for j := i + 1; j < n; j++ {
			cur += a[j]
			if cur == pre {
				cnt++
				cur = 0
			} else if cur > pre {
				break
			}
		}

		if cur == 0 {
			ans = min(ans, n-cnt)
		}
	}

	fmt.Fprintln(out, ans)
}

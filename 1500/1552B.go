package main

import (
	"bufio"
	"fmt"
	"os"
)

// codeforces 1552B
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
	r := make([][5]int, n+1)
	for i := 1; i <= n; i++ {
		for j := range 5 {
			fmt.Fscan(in, &r[i][j])
		}
	}

	ans := 1
	for i := 2; i <= n; i++ {
		cnt := 0
		for j := range 5 {
			if r[ans][j] < r[i][j] {
				cnt++
			}
		}

		if cnt < 3 {
			ans = i
		}
	}

	for i := 1; i <= n; i++ {
		if i == ans {
			continue
		}

		cnt := 0
		for j := range 5 {
			if r[ans][j] < r[i][j] {
				cnt++
			}
		}

		if cnt < 3 {
			fmt.Fprintln(out, -1)
			return
		}
	}

	fmt.Fprintln(out, ans)
}

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
	var n, a, b int
	fmt.Fscan(in, &n, &a, &b)
	if a+b+2 > n || abs(a-b) > 1 {
		fmt.Fprintln(out, -1)
		return
	}
	cnt := a + b + 2
	pos := 1
	ans := make([]int, n)
	if a == b {
		for i := 0; i < cnt; i += 2 {
			ans[i] = pos
			pos++
		}
		for i := 1; i < cnt; i += 2 {
			ans[i] = pos
			pos++
		}
		for i := cnt; i < n; i++ {
			ans[i] = pos
			pos++
		}
	} else if a > b {
		for i := n - 1; i >= cnt; i-- {
			ans[i] = pos
			pos++
		}
		for i := 0; i < cnt; i += 2 {
			ans[i] = pos
			pos++
		}
		for i := 1; i < cnt; i += 2 {
			ans[i] = pos
			pos++
		}
	} else {
		cnt = n - cnt
		for i := cnt + 1; i < n; i += 2 {
			ans[i] = pos
			pos++
		}
		for i := cnt + 2; i < n; i += 2 {
			ans[i] = pos
			pos++
		}
		for i := cnt; i >= 0; i-- {
			ans[i] = pos
			pos++
		}
	}
	for _, x := range ans {
		fmt.Fprint(out, x, " ")
	}
	fmt.Fprintln(out)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

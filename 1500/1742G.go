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

	ans := make([]int, 0, n)
	cur := 0
	used := make([]bool, n)
	for {
		idx, val := -1, cur
		for i := range n {
			if !used[i] && (cur|a[i]) > val {
				idx, val = i, cur|a[i]
			}
		}

		if idx == -1 {
			break
		}
		ans = append(ans, a[idx])
		cur = val
		used[idx] = true
	}
	for i := range n {
		if !used[i] {
			ans = append(ans, a[i])
		}
	}
	for i := range ans {
		fmt.Fprint(out, ans[i], " ")
	}
	fmt.Fprintln(out)
}

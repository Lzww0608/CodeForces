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

	var n, a, b int
	fmt.Fscan(in, &n, &a, &b)
	if min(a, b) > 1 {
		fmt.Fprintln(out, "NO")
		return
	}

	ans := make([][]int, n)
	for i := 0; i < n; i++ {
		ans[i] = make([]int, n)
	}

	if a == 1 && b == 1 {
		if n == 2 || n == 3 {
			fmt.Fprintln(out, "NO")
			return
		}
		fmt.Fprintln(out, "YES")
		for i := 1; i < n; i++ {
			ans[i][i-1] = 1
			ans[i-1][i] = 1
		}
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				fmt.Fprint(out, ans[i][j])
			}
			fmt.Fprintln(out)
		}
		return
	}

	f := true
	if a == 1 {
		a, b = b, a
		f = false
	}

	for i := n - a; i > 0; i-- {
		ans[i][i-1] = 1
		ans[i-1][i] = 1
	}

	fmt.Fprintln(out, "YES")
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == j {
				fmt.Fprint(out, 0)
				continue
			}
			if f {
				fmt.Fprint(out, ans[i][j])
			} else {
				fmt.Fprint(out, ans[i][j]^1)
			}

		}
		fmt.Fprintln(out)
	}
}

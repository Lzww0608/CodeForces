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

	var s string
	var n int
	fmt.Fscan(in, &s, &n)

	ans := make([]int, n)
	var f func(int, int) bool
	f = func(p, d int) bool {
		if p == n {
			return true
		}
		for ans[p] = d; ans[p] < 10; ans[p]++ {
			if s[ans[p]] == '1' && (p == 0 || ans[p] != ans[p-1]) && f(p+1, ans[p]-d+1) {
				return true
			}
		}
		return false
	}

	if f(0, 0) {
		fmt.Fprintln(out, "YES")
		for _, x := range ans {
			fmt.Fprint(out, x+1, " ")
		}
	} else {
		fmt.Fprintln(out, "NO")
	}

}

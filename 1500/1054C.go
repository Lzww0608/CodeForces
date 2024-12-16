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

	var n int
	fmt.Fscan(in, &n)
	l := make([]int, n)
	r := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &l[i])
	}
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &r[i])
	}

	if l[0] != 0 || r[n-1] != 0 {
		fmt.Fprintln(out, "NO")
		return
	}

	ans := make([]int, n)
	for i := 0; i < n; i++ {
		ans[i] = n - l[i] - r[i]
	}
	for k, x := range ans {
		cnt := 0
		for i := 0; i < k; i++ {
			if ans[i] > x {
				cnt++
			}
		}
		if cnt != l[k] {
			fmt.Fprintln(out, "NO")
			return
		}
		cnt = 0
		for i := k + 1; i < n; i++ {
			if ans[i] > x {
				cnt++
			}
		}
		if cnt != r[k] {
			fmt.Fprintln(out, "NO")
			return
		}
	}

	fmt.Fprintln(out, "YES")
	for _, x := range ans {
		fmt.Fprint(out, x, " ")
	}
}

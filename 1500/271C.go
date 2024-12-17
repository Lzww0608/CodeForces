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

	var n, k int
	fmt.Fscan(in, &n, &k)
	if k*3 > n {
		fmt.Fprintln(out, -1)
		return
	}
	ans := make([]int, n)
	for i := 0; i < n && k > 1; i += 2 {
		ans[i], ans[i+1] = k, k
		ans[n-k] = k
		k--
	}

	for i := 0; i < n; i++ {
		if ans[i] == 0 {
			ans[i] = 1
		}
	}
	for _, x := range ans {
		fmt.Fprint(out, x, " ")
	}
}

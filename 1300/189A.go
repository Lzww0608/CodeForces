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

	var n, a, b, c int
	fmt.Fscan(in, &n, &a, &b, &c)
	f := make([]int, n+1)
	for i := 0; i < n; i++ {
		if i > 0 && f[i] == 0 {
			continue
		}
		for _, x := range []int{a, b, c} {
			if i+x <= n {
				f[i+x] = max(f[i+x], f[i]+1)
			}
		}
	}

	fmt.Fprintln(out, f[n])
}

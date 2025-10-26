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
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
	}

	ans := 0
	f := make([]int, n)
	for i := 0; i < n; i++ {
		f[i] = 1
		if i > 0 && a[i] > a[i-1] {
			f[i] = f[i-1] + 1
		}

		ans = max(ans, f[i])
		if i-f[i]-1 >= 0 && a[i-f[i]+1] > a[i-f[i]-1] {
			ans = max(ans, f[i]+f[i-f[i]-1])
		} else if i-f[i] >= 0 && i-f[i]+2 < n && a[i-f[i]+2] > a[i-f[i]] {
			ans = max(ans, f[i]+f[i-f[i]]-1)
		}
	}

	fmt.Fprintln(out, ans)
}

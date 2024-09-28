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

	var t, n, k, x int
	for fmt.Fscan(in, &t); t > 0; t-- {
		fmt.Fscan(in, &n, &k, &x)
		if k > n || k > x+1 {
			fmt.Fprintln(out, -1)
			continue
		}
		ans := k * (k - 1) / 2
		if k >= x {
			ans += (k - 1) * (n - k)
		} else {
			ans += x * (n - k)
		}

		fmt.Fprintln(out, ans)
	}
}

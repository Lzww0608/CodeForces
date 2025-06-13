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
	var n, k int
	fmt.Fscan(in, &n, &k)
	if k >= n {
		fmt.Fprintln(out, 1)
		return
	}
	ans := n
	for i := 2; i <= k && i*i <= n; i++ {
		if n%i == 0 {
			if n/i < ans {
				ans = n / i
			}

			if n/i <= k && i < ans {
				ans = i
			}
		}
	}
	fmt.Fprintln(out, ans)
}

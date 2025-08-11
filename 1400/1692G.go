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
	a := make([]int, n)
	f := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
	}

	ans := 0
	f[0] = 1
	for i := 1; i < n; i++ {
		if a[i] > a[i-1]/2 {
			f[i] = f[i-1] + 1
			if f[i] >= k+1 {
				ans++
			}
		} else {
			f[i] = 1
		}
	}

	fmt.Fprintln(out, ans)
}

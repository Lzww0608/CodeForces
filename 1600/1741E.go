package main

import (
	"bufio"
	"fmt"
	"os"
)

// codeforces 1741E

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
	a := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i])
	}

	f := make([]bool, n+1)
	f[0] = true
	for i := 1; i <= n; i++ {
		if f[i-1] && i+a[i] <= n {
			f[i+a[i]] = true
		}
		if i-a[i]-1 >= 0 && f[i-a[i]-1] {
			f[i] = true
		}
	}

	if f[n] {
		fmt.Fprintln(out, "YES")
	} else {
		fmt.Fprintln(out, "NO")
	}
}

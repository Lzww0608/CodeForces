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
	var n int
	fmt.Fscan(in, &n)
	if n&1 == 0 {
		fmt.Fprintln(out, "NO")
		return
	}
	fmt.Fprintln(out, "YES")
	for i := 1; i <= n/2+1; i++ {
		fmt.Fprintln(out, i, i+n+n/2)
	}
	for i := n; i > n/2+1; i-- {
		fmt.Fprintln(out, i, n+n/2-(n-i))
	}
}

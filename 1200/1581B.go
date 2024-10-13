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
		sovle(in, out)
	}
}

func sovle(in *bufio.Reader, out *bufio.Writer) {
	var n, m, k int
	fmt.Fscan(in, &n, &m, &k)
	if m < n-1 || m > n*(n-1)/2 {
		fmt.Fprintln(out, "NO")
		return
	}
	d := 0
	if n > 1 {
		if m == n*(n-1)/2 {
			d = 1
		} else {
			d = 2
		}
	}
	if d < k-1 {
		fmt.Fprintln(out, "YES")
	} else {
		fmt.Fprintln(out, "NO")
	}
}

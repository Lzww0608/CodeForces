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

	var n, m, l, r int
	fmt.Fscan(in, &n, &m)
	mn := n
	for i := 0; i < m; i++ {
		fmt.Fscan(in, &l, &r)
		if r-l < mn {
			mn = r - l
		}
	}

	mn++
	fmt.Fprintln(out, mn)
	for i := 0; i < n; i++ {
		fmt.Fprint(out, i%mn, " ")
	}
}

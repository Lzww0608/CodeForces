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

	var n, m, k int
	fmt.Fscan(in, &n, &m, &k)
	var s string
	fmt.Fscan(in, &s)
	if s[m-1] == s[k-1] {
		fmt.Fprintln(out, 0)
		return
	}
	fmt.Fprintln(out, 1)
	return
}

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
	var n, m int
	fmt.Fscan(in, &n, &m)
	ok := false
	a := make([][4]int, n)
	for i := range a {
		for j := range a[i] {
			fmt.Fscan(in, &a[i][j])
		}
		if a[i][1] == a[i][2] {
			ok = true
		}
	}

	if m&1 == 1 || !ok {
		fmt.Fprintln(out, "NO")
		return
	}

	fmt.Fprintln(out, "YES")
	return
}

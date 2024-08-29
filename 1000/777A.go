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

	var n, x int
	fmt.Fscan(in, &n)
	fmt.Fscan(in, &x)
	n = (n - 1) % 6
	a := [][]int{{1, 0, 2}, {1, 2, 0}, {2, 1, 0}, {2, 0, 1}, {0, 2, 1}, {0, 1, 2}}
	fmt.Fprintln(out, a[n][x])
}

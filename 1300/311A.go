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

	var n, k int
	fmt.Fscan(in, &n, &k)
	if k >= n*(n-1)/2 {
		fmt.Fprintln(out, "no solution")
		return
	}
	for i := 0; i < n; i++ {
		fmt.Fprintln(out, 0, i)
	}
}

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

	var n int
	fmt.Fscan(in, &n)

	if n <= 2 {
		fmt.Fprintln(out, -1)
		return
	}

	if n&1 == 1 {
		m := (n*n - 1) / 2
		k := (n*n + 1) / 2
		fmt.Fprintln(out, m, k)
	} else {
		// k - m = 2
		// k + m = n * n / 2
		// k = n * n / 4 + 2
		k := n*n/4 + 1
		m := k - 2
		fmt.Fprintln(out, m, k)
	}
}

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

	var n, x, sum int
	fmt.Fscan(in, &n)
	mx := 0
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &x)
		sum += x
		mx = max(mx, x)
	}
	// (x - a1) + (x - a2) + (x - a3) + ... >= x
	// (n - 1) * x >= a1 + ... + an
	// x >= (a1 + ... + an) / (n - 1)
	fmt.Fprintln(out, max((sum+n-2)/(n-1), mx))
}

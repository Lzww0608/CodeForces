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

	var t, n, k int
	for fmt.Fscan(in, &t); t > 0; t-- {
		fmt.Fscan(in, &n, &k)
		if k >= n {
			fmt.Fprintln(out, k-n)
		} else if k&1 == n&1 {
			fmt.Fprintln(out, 0)
		} else {
			fmt.Fprintln(out, 1)
		}

	}

}

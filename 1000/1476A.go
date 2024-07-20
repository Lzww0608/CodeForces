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
		if n%k == 0 {
			fmt.Fprintln(out, 1)
		} else if n < k {
			fmt.Fprintln(out, (k+n-1)/n)
		} else {
			fmt.Fprintln(out, 2)
		}
	}

}

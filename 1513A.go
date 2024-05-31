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
		if k > (n-1)/2 {
			fmt.Fprintln(out, -1)
			continue
		}
		for i := 0; i < k; i++ {
			fmt.Fprint(out, i+1, " ", n-i, " ")
		}
		for i := k + 1; i <= n-k; i++ {
			fmt.Fprint(out, i, " ")
		}
		fmt.Fprintln(out)
	}
}

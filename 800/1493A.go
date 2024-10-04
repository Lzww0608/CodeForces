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
		ans := max(0, n-k)
		ans += max(0, k>>1)
		fmt.Fprintln(out, ans)
		for i := n; i > k; i-- {
			fmt.Fprint(out, i, " ")
		}
		for i := k - 1; i >= (k+1)/2; i-- {
			fmt.Fprint(out, i, " ")
		}
		fmt.Fprintln(out)
	}

}

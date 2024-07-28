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
		cnt := 1
		for i := 1; cnt <= n; cnt++ {
			fmt.Fprint(out, i, " ")
			if i+cnt > k-(n-cnt-1) {
				i++
			} else {
				i += cnt
			}
		}
		fmt.Fprintln(out)
	}
}

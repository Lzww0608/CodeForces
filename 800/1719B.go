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
		if k%4 == 0 {
			fmt.Fprintln(out, "NO")
			continue
		}
		fmt.Fprintln(out, "YES")
		for i := 1; i <= n; i += 2 {
			if (i+k)*(i+1)%4 == 0 {
				fmt.Fprintln(out, i, i+1)
			} else {
				fmt.Fprintln(out, i+1, i)
			}
		}

	}
}

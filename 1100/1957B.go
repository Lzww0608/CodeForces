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
		t := 1
		for t <= k {
			t <<= 1
		}
		t >>= 1
		t -= 1
		if n == 1 {
			fmt.Fprintln(out, k)
			continue
		}

		fmt.Fprint(out, t, " ", k-t)
		for i := 3; i <= n; i++ {
			fmt.Fprint(out, " ", 0)
		}
		fmt.Fprintln(out)
	}
}

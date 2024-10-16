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
	t, cur := n-1, 1
	for i := 1; i <= n; i++ {
		if k == 1 {
			if t < 0 {
				t = -1
			} else {
				t = 1
			}
		}

		if k > 1 {
			k--
			fmt.Fprint(out, cur, " ")
			cur += t
			if t > 0 {
				t = -(t - 1)
			} else {
				t = -(t + 1)
			}
		} else {
			fmt.Fprint(out, cur, " ")
			cur += t
			k = 0
		}

	}
	fmt.Fprintln(out)
}

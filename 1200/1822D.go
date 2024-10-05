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

	var t, n int
	for fmt.Fscan(in, &t); t > 0; t-- {
		fmt.Fscan(in, &n)
		if n != 1 && n&1 == 1 {
			fmt.Fprintln(out, -1)
			continue
		}
		cur := n
		target := 0
		fmt.Fprint(out, cur, " ")
		for i := 1; i < n; i++ {
			if i&1 == 1 {
				target += n - i
			} else {
				target -= n - i
			}
			if cur%n < target {
				fmt.Fprint(out, target-cur%n, " ")
			} else {
				fmt.Fprint(out, target+n-cur%n, " ")
			}
			cur = target
		}
		fmt.Fprintln(out)
	}
}

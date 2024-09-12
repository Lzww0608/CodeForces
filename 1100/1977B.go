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
		ans := make([]int, 32)
		for i := range ans {
			if (n>>i)&1 == 1 {
				if ans[i] == 1 {
					ans[i+1] = 1
					ans[i] = 0
				} else if i > 0 {
					if ans[i-1] == 1 {
						ans[i+1] = 1
						ans[i-1] = -1
					} else {
						ans[i] = 1
					}
				} else if i == 0 {
					ans[i] = 1
				}
			}
		}
		fmt.Fprintln(out, 32)
		for _, x := range ans {
			fmt.Fprint(out, x, " ")
		}
		fmt.Fprintln(out)
	}
}

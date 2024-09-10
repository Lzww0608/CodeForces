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
		if n == 1 {
			fmt.Fprintln(out, 0)
			continue
		}
		ans := n
		for i := 1; i < n; i++ {
			cur := i + (n+i)/(i+1) - 1
			//fmt.Fprintln(out, i, cur)
			if cur <= ans {
				ans = cur
			} else {
				break
			}
		}
		fmt.Fprintln(out, ans)
	}
}

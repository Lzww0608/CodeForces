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
		a := make([]int, n)
		cnt := 0
		for i := range a {
			fmt.Fscan(in, &a[i])
			if a[i] == 1 {
				cnt++
			}
		}
		zero := 0
		ans := cnt
		for i := n - 1; i >= 0; i-- {
			if a[i] == 0 {
				zero++
			} else {
				cnt--
			}
			ans = min(ans, max(zero, cnt))
		}
		fmt.Fprintln(out, ans)
	}
}

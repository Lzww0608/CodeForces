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
		ans, cnt := 0, 0
		for i := range a {
			fmt.Fscan(in, &a[i])
			if i < n-1 {
				ans += a[i]
				if a[i] == 0 {
					cnt++
				}

			}
		}
		for j := 0; j < n-1 && a[j] == 0; j++ {
			cnt--
		}

		fmt.Fprintln(out, ans+cnt)
	}

}

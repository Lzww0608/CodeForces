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
		for i := range a {
			fmt.Fscan(in, &a[i])
		}
		cnt := 2
		for i := 1; i < n; i++ {
			if a[i] >= a[i-1] {
				cnt++
			} else {
				break
			}
		}
		for i := n - 2; i >= 0; i-- {
			if a[i] >= a[i+1] {
				cnt++
			} else {
				break
			}
		}

		if cnt >= n {
			fmt.Fprintln(out, "YES")
		} else {
			fmt.Fprintln(out, "NO")
		}
	}
}

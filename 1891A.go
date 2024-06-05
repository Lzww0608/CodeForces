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
next:
	for fmt.Fscan(in, &t); t > 0; t-- {
		fmt.Fscan(in, &n)
		a := make([]int, n)
		for i := range a {
			fmt.Fscan(in, &a[i])
		}
		for i := 2; n > (1 << (i - 1)); i++ {
			for j := 1 << (i - 1); j < n-1 && j < (1<<i)-1; j++ {
				if a[j] > a[j+1] {
					fmt.Fprintln(out, "NO")
					continue next
				}
			}
		}
		fmt.Fprintln(out, "YES")
	}

}

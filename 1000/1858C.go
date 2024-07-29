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
		a := make([]int, n+1)
		for i := 1; i <= n; i++ {
			if a[i] == 1 {
				continue
			}
			for j := i; j <= n; j *= 2 {
				if a[j] == 1 {
					break
				}
				a[j] = 1
				fmt.Fprint(out, j, " ")
			}

		}
		fmt.Fprintln(out)

	}

}

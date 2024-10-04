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

	var t, n, x int
	for fmt.Fscan(in, &t); t > 0; t-- {
		fmt.Fscan(in, &n, &x)
		a := make([]int, n)
		sum := 0
		for i := range a {
			fmt.Fscan(in, &a[i])
			sum += a[i]
		}
		if sum == x {
			fmt.Fprintln(out, "NO")
			continue
		}
		fmt.Fprintln(out, "YES")
		sum = 0
		for i := 0; i < n; i++ {
			if sum+a[i] == x {
				a[i], a[i+1] = a[i+1], a[i]
			}
			sum += a[i]
			fmt.Fprint(out, a[i], " ")
		}
		fmt.Fprintln(out)

	}
}

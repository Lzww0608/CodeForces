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

	var n, k, t int
	for fmt.Fscan(in, &t); t > 0; t-- {
		fmt.Fscan(in, &n, &k)
		a := make([]int, n)
		sum := 0
		for i := range a {
			fmt.Fscan(in, &a[i])
			sum += a[i]
		}
		if k >= sum {
			fmt.Fprintln(out, n)
			continue
		}
		x, y := (k+1)/2, k/2
		ans := 0
		for i := 0; i < n; i++ {
			if x >= a[i] {
				x -= a[i]
				a[i] = 0
				ans++
			} else {
				a[i] -= x
				break
			}
		}

		for i := n - 1; i >= 0 && a[i] > 0; i-- {
			if y >= a[i] {
				y -= a[i]
				a[i] = 0
				ans++
			} else {
				break
			}

		}
		fmt.Fprintln(out, ans)
	}
}

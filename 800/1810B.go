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
		if n&1 == 0 {
			fmt.Fprintln(out, -1)
			continue
		}
		a := []int{}
		for n > 1 {
			if (n+1)/2%2 == 1 {
				a = append(a, 1)
				n = (n + 1) / 2
			} else {
				a = append(a, 2)
				n = (n - 1) / 2
			}
		}
		fmt.Fprintln(out, len(a))
		for i := len(a) - 1; i >= 0; i-- {
			fmt.Fprint(out, a[i], " ")
		}
		fmt.Fprintln(out)

	}
}

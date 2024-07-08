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
		b := make([]int, n)
		minA, minB := int(1e9), int(1e9)
		sumA, sumB := 0, 0
		for i := range a {
			fmt.Fscan(in, &a[i])
			sumA += a[i]
			minA = min(minA, a[i])
		}
		for i := range b {
			fmt.Fscan(in, &b[i])
			minB = min(minB, b[i])
			sumB += b[i]
		}

		fmt.Fprintln(out, min(sumA+n*minB, sumB+n*minA))
	}
}

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

	var t int
	for fmt.Fscan(in, &t); t > 0; t-- {
		solve(in, out)
	}
}

func solve(in *bufio.Reader, out *bufio.Writer) {
	var q, d, k int
	fmt.Fscan(in, &q)
	var s string
	fA, fB := false, false
	sumA, sumB := 0, 0

	for i := 0; i < q; i++ {
		fmt.Fscan(in, &d, &k, &s)
		if d == 1 {
			for _, c := range s {
				if c != 'a' {
					fA = true
				} else {
					sumA += k
				}
			}
		} else {
			for _, c := range s {
				if c != 'a' {
					fB = true
				} else {
					sumB += k
				}
			}
		}
		if fB || !fA && sumA < sumB {
			fmt.Fprintln(out, "YES")
		} else {
			fmt.Fprintln(out, "NO")
		}

	}
}

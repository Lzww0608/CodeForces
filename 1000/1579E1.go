package main

import (
	"bufio"
	"fmt"
	"math"
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
		b := make([]int, n*2)
		l, r := n-1, n
		b[r] = math.MaxInt32
		for _, x := range a {
			if x < b[l+1] {
				b[l] = x
				l--
			} else {
				b[r] = x
				r++
			}
		}
		for i := l + 1; i < r; i++ {
			fmt.Fprint(out, b[i], " ")
		}
		fmt.Fprintln(out)
	}
}

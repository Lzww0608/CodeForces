package main

import (
	"bufio"
	. "fmt"
	"os"
)

func main() {

	var q int
	in := bufio.NewReader(os.Stdin)
	Fscan(in, &q)

	var n, k, t int
	for ; q > 0; q-- {

		Fscan(in, &n, &k)

		a := make([]int, k)
		f := make([]int, n)
		for i := range f {
			f[i] = 2e9
		}

		for i := range a {
			Fscan(in, &a[i])
		}

		for i := range a {
			Fscan(in, &t)
			f[a[i]-1] = t
		}

		for i := n - 2; i >= 0; i-- {
			f[i] = min(f[i], f[i+1]+1)
		}

		mn := int(2e9)
		for _, x := range f {
			mn = min(mn+1, x)
			Print(mn, " ")
		}
		Println()
	}
	return
}

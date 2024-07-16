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

	var n int
	fmt.Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		fmt.Fscan(in, &a[i])
	}

	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		if i&1 == 0 {
			a[i], a[j] = a[j], a[i]
		}
	}

	for _, x := range a {
		fmt.Fprint(out, x, " ")
	}
	fmt.Fprintln(out)
}

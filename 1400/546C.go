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

	var n, k1, k2 int
	fmt.Fscan(in, &n)
	fmt.Fscan(in, &k1)
	a := make([]int, k1)
	for i := 0; i < k1; i++ {
		fmt.Fscan(in, &a[i])
	}
	fmt.Fscan(in, &k2)
	b := make([]int, k2)
	for i := 0; i < k2; i++ {
		fmt.Fscan(in, &b[i])
	}

	for cnt := 1; cnt < n*n*n*n*n; cnt++ {
		x, y := a[0], b[0]
		a = a[1:]
		b = b[1:]
		if x < y {
			b = append(b, x)
			b = append(b, y)
		} else {
			a = append(a, y)
			a = append(a, x)
		}

		if len(a) == 0 {
			fmt.Fprintln(out, cnt, 2)
			return
		} else if len(b) == 0 {
			fmt.Fprintln(out, cnt, 1)
			return
		}
	}

	fmt.Fprintln(out, -1)
}

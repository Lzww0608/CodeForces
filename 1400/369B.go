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

	var n, k, l, r, sA, sK int
	fmt.Fscan(in, &n, &k, &l, &r, &sA, &sK)
	sA -= sK
	t := r
	for i := k; i > 0; i-- {
		for i*t > sK && i*(t-1) >= sK {
			t--
		}
		fmt.Fprint(out, t, " ")
		sK -= t
	}

	for i := n - k; i > 0; i-- {
		for i*t > sA && i*(t-1) >= sA {
			t--
		}
		fmt.Fprint(out, t, " ")
		sA -= t
	}
}

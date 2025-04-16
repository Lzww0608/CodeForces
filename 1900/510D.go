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
	b := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
	}
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &b[i])
	}
	f := make(map[int]int)
	f[0] = 0
	for i, x := range a {
		for y := range f {
			if g := gcd(x, y); f[g] == 0 || f[g] > b[i]+f[y] {
				f[g] = b[i] + f[y]
			}
		}
	}

	if f[1] > 0 {
		fmt.Fprintln(out, f[1])
	} else {
		fmt.Fprintln(out, -1)
	}
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

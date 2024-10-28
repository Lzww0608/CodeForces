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
	var x int
	fmt.Fscan(in, &x)
	a := []int{x}
	for k := 1; x > 2; {
		if x&k != 0 {
			if x > k {
				x -= k
				k <<= 1
			} else {
				x -= k / 2
				k /= 2
			}
			a = append(a, x)
		} else {
			k <<= 1
		}
	}
	a = append(a, 1)
	fmt.Fprintln(out, len(a))
	for _, x := range a {
		fmt.Fprint(out, x, " ")
	}
	fmt.Fprintln(out)
}

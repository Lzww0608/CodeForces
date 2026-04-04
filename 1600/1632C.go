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
	var a, b int
	fmt.Fscan(in, &a, &b)
	ans := b - a

	// (a' - a) + (b' - b) + ((a' | b') - b') + 1
	// a' + (a' | b') + 1 - a - b
	for x := a; x < b; x++ {
		cur := x + 1 - a - b
		y := 0
		for i := 21; i >= 0; i-- {
			if (b & (1 << i)) != 0 {
				y += (1 << i)
			} else if (x & (1 << i)) != 0 {
				y += (1 << i)
				break
			}
		}

		ans = min(ans, cur+(x|y))
	}

	fmt.Fprintln(out, ans)
}

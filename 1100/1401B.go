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
	var x1, x2, y1, y2, z1, z2 int
	for fmt.Fscan(in, &t); t > 0; t-- {
		fmt.Fscan(in, &x1, &y1, &z1)
		fmt.Fscan(in, &x2, &y2, &z2)
		ans := min(z1, y2) * 2
		y2 -= min(y2, z1)
		y1 -= x2 + y2
		if y1 > 0 {
			ans -= 2 * y1
		}
		fmt.Fprintln(out, ans)
	}
}

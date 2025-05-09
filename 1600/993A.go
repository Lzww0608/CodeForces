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

	var x1, y1, x2, y2, x3, y3, x4, y4 int
	fmt.Fscan(in, &x1, &y1, &x2, &y2, &x3, &y3, &x4, &y4)
	xMin := min(x1, x2, x3, x4)
	xMax := max(x1, x2, x3, x4)
	yMin := min(y1, y2, y3, y4)
	yMax := max(y1, y2, y3, y4)

	var a1, b1, a2, b2, a3, b3, a4, b4 int
	fmt.Fscan(in, &a1, &b1, &a2, &b2, &a3, &b3, &a4, &b4)
	uMin := min(a1+b1, a2+b2, a3+b3, a4+b4)
	uMax := max(a1+b1, a2+b2, a3+b3, a4+b4)
	vMin := min(b1-a1, b2-a2, b3-a3, b4-a4)
	vMax := max(b1-a1, b2-a2, b3-a3, b4-a4)

	check := func(i, j int) bool {
		i, j = i+j, j-i
		if i >= uMin && i <= uMax && j >= vMin && j <= vMax {
			return true
		}
		return false
	}

	for i := xMin; i <= xMax; i++ {
		for j := yMin; j <= yMax; j++ {
			if check(i, j) {
				fmt.Fprintln(out, "YES")
				return
			}
		}
	}
	fmt.Fprintln(out, "NO")
	return
}

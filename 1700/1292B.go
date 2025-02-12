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

	var x0, y0, ax, ay, bx, by, xs, ys, t int
	fmt.Fscan(in, &x0, &y0, &ax, &ay, &bx, &by, &xs, &ys, &t)
	v := [][2]int{}
	x, y := x0, y0
	for x <= 2e16 && y <= 2e16 {
		v = append(v, [2]int{x, y})
		x = x*ax + bx
		y = y*ay + by
	}

	ans := 0
	for i := 0; i < len(v); i++ {
		for j := 0; j < len(v); j++ {
			if abs(v[i][0]-xs)+abs(v[i][1]-ys)+abs(v[i][0]-v[j][0])+abs(v[i][1]-v[j][1]) <= t {
				ans = max(ans, abs(j-i)+1)
			}
		}
	}

	fmt.Fprintln(out, ans)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

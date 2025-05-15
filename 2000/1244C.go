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

	var n, p, w, d int
	fmt.Fscan(in, &n, &p, &w, &d)
	x, y, z := 0, 0, 0
	for y < w && (p-y*d)%w != 0 {
		y++
	}
	if y == w {
		fmt.Fprintln(out, -1)
		return
	}
	x = (p - y*d) / w
	if x+y > n || x < 0 {
		fmt.Fprintln(out, -1)
		return
	}
	z = n - x - y
	fmt.Fprintln(out, x, y, z)
}

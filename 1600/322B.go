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

	var r, g, b int
	fmt.Fscan(in, &r, &g, &b)
	ans := r/3 + g/3 + b/3
	x, y, z := r%3, g%3, b%3
	cur := min(x, y, z)
	if x == 0 && r >= 3 {
		cur = max(cur, min(x+3, y, z)-1)
	} else if y == 0 && g >= 3 {
		cur = max(cur, min(x, y+3, z)-1)
	} else if z == 0 && b >= 3 {
		cur = max(cur, min(x, y, z+3)-1)
	}

	ans += cur
	fmt.Fprintln(out, ans)
}

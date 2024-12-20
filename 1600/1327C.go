package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, m, k int
	fmt.Fscan(in, &n, &m, &k)
	x, y := 0, 0
	a := make([][2]int, k)
	b := make([][2]int, k)
	for i := 0; i < k; i++ {
		fmt.Fscan(in, &a[i][0], &a[i][1])
		x = max(x, a[i][0])
		y = max(y, a[i][1])
	}
	for i := 0; i < k; i++ {
		fmt.Fscan(in, &b[i][0], &b[i][1])
	}

	cnt := x + y - 2 + m*n
	fmt.Fprintln(out, cnt)
	fmt.Fprint(out, strings.Repeat("U", x-1))
	fmt.Fprint(out, strings.Repeat("L", y-1))
	for i := 0; i < n; i++ {
		if i&1 == 0 {
			fmt.Fprint(out, strings.Repeat("R", m-1))
		} else {
			fmt.Fprint(out, strings.Repeat("L", m-1))
		}
		fmt.Fprint(out, strings.Repeat("D", 1))
	}
	fmt.Fprintln(out)
}

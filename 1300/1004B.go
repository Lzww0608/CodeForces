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

	var n, m, x, y int
	fmt.Fscan(in, &n, &m)
	for i := 0; i < m; i++ {
		fmt.Fscan(in, &x, &y)
	}
	for i := 0; i < n; i++ {
		if i&1 == 0 {
			fmt.Fprint(out, "0")
		} else {
			fmt.Fprint(out, "1")
		}
	}
}

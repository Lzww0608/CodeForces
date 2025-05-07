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

	var n, m, k int
	fmt.Fscan(in, &n, &m, &k)

	var x, y int
	for i := 0; i < k; i++ {
		fmt.Fscan(in, &x, &y)
		if x <= 5 || y <= 5 || x+4 >= n || y+4 >= m {
			fmt.Fprintln(out, "YES")
			return
		}
	}
	fmt.Fprintln(out, "NO")
}

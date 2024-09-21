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

	var n, x int
	fmt.Fscan(in, &n, &x)
	a := make([][2]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i][0], &a[i][1])
	}
	x = 7 - x
	for i := 0; i < n; i++ {
		if a[i][0] == x || a[i][1] == x || a[i][0] == 7-x || a[i][1] == 7-x {
			fmt.Fprintln(out, "NO")
			return
		}
	}
	fmt.Fprintln(out, "YES")

}

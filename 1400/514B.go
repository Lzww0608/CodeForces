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

	var n, x, y int
	fmt.Fscan(in, &n, &x, &y)
	ans := 0
	a := make([][2]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i][0], &a[i][1])
	}

	for i := 0; i < n; i++ {
		ans++
		u := a[i]
		for j := i - 1; j >= 0; j-- {
			v := a[j]
			if (u[1]-y)*(v[0]-x) == (u[0]-x)*(v[1]-y) {
				ans--
				break
			}
		}
	}

	fmt.Fprintln(out, ans)
}

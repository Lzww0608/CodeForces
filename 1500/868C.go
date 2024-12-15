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

	var n, k int
	fmt.Fscan(in, &n, &k)
	a := make([][]int, n)
	has := make([]bool, 1<<k)
	for i := range a {
		a[i] = make([]int, k)
		v := 0
		for j := range a[i] {
			fmt.Fscan(in, &a[i][j])
			v |= a[i][j] << j
		}
		has[v] = true
	}
	if has[0] {
		fmt.Fprintln(out, "YES")
		return
	}
	for i, x := range has {
		for j, y := range has {
			if i != j && x && y && i&j == 0 {
				fmt.Fprintln(out, "YES")
				return
			}
		}
	}
	fmt.Fprintln(out, "NO")

}

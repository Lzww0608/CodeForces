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

	var n int
	fmt.Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		fmt.Fscan(in, &a[i])
	}

	b := [2]int{}
	ans := 0
	for i, x := range a {
		y := x & 1
		if b[y] != 0 {
			ans = 1 - y
		} else {
			b[y] = i + 1
		}
	}

	fmt.Fprintln(out, b[ans])
}

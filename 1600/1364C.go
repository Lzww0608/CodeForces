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
	b := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
		b[i] = -1
	}

	exist := make(map[int]bool)
	for i := 1; i < n; i++ {
		if a[i] != a[i-1] {
			b[i] = a[i-1]
			exist[b[i]] = true
		}
	}

	exist[a[n-1]] = true
	x := 0
	for i := 0; i < n; i++ {
		if b[i] != -1 {
			continue
		}
		for exist[x] {
			x++
		}
		b[i] = x
		x++
	}
	for _, x := range b {
		fmt.Fprint(out, x, " ")
	}
}

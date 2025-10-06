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
	fmt.Fscan(in, &n)
	a := make([]int, n+1)
	b := make([]int, n+1)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
	}
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &x)
		b[x] = i
	}

	cur := -1
	ans := 0
	for i := range n {
		x = a[i]
		pos := b[x]
		if cur > pos {
			ans++
		} else {
			cur = max(cur, pos)
		}
	}

	fmt.Fprintln(out, ans)
}

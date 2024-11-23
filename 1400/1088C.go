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
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
	}
	fmt.Fprintln(out, n+1)
	sum := 0
	for i := n - 1; i >= 0; i-- {
		cur := (i - (a[i]+sum)%n + n) % n
		fmt.Fprintln(out, 1, i+1, cur)
		sum += cur
	}
	fmt.Fprintln(out, 2, n, n)
}

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
	}
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &b[i])
	}

	f := make([][2]int, n+1)
	for i := 0; i < n; i++ {
		f[i+1][0] = max(f[i][1]+a[i], f[i][0])
		f[i+1][1] = max(f[i][0]+b[i], f[i][1])
	}

	fmt.Fprintln(out, max(f[n][0], f[n][1]))
}

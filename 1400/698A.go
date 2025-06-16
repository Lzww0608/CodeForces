package main

import (
	"bufio"
	"fmt"
	"math"
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
	f := make([][3]int, n+1)

	for i, x := range a {
		if x == 0 {
			f[i+1][0] = min(f[i][0], f[i][1], f[i][2]) + 1
			f[i+1][1] = math.MaxInt32
			f[i+1][2] = math.MaxInt32
		} else if x == 1 {
			f[i+1][0] = min(f[i][0], f[i][1], f[i][2]) + 1
			f[i+1][1] = math.MaxInt32
			f[i+1][2] = min(f[i][0], f[i][1])
		} else if x == 2 {
			f[i+1][0] = min(f[i][0], f[i][1], f[i][2]) + 1
			f[i+1][2] = math.MaxInt32
			f[i+1][1] = min(f[i][0], f[i][2])
		} else {
			f[i+1][0] = min(f[i][0], f[i][1], f[i][2]) + 1
			f[i+1][1] = min(f[i][0], f[i][2])
			f[i+1][2] = min(f[i][0], f[i][1])
		}
	}

	fmt.Fprintln(out, min(f[n][0], f[n][1], f[n][2]))
}

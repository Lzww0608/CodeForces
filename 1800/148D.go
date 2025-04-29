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

	var w, b int
	fmt.Fscan(in, &w, &b)
	f := make([][]float64, w+1)
	for i := range f {
		f[i] = make([]float64, b+1)
		f[i][0] = 1.0
	}
	if w == 0 {
		fmt.Fprintln(out, 0.0)
		return
	}

	for i := 1; i <= w; i++ {
		for j := 1; j <= b; j++ {
			f[i][j] = float64(i) / float64(i+j)
			if j >= 2 {
				f[i][j] += float64(j) / float64(i+j) * float64(j-1) / float64(i+j-1) * float64(i) / float64(i+j-2) * f[i-1][j-2]
			}
			if j >= 3 {
				f[i][j] += float64(j) / float64(i+j) * float64(j-1) / float64(i+j-1) * float64(j-2) / float64(i+j-2) * f[i][j-3]
			}
		}
	}

	fmt.Fprintln(out, f[w][b])
}

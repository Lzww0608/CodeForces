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

	var t, h, w int
	fmt.Fscan(in, &t)
	for ; t > 0; t-- {
		fmt.Fscan(in, &h, &w)
		a := make([][]int, h+2)
		for i := range a {
			a[i] = make([]int, w+2)
		}

		for i := 1; i <= h; i++ {
			if i == 1 || i == h {
				for j := 1; j <= w; j++ {
					if a[i-1][j] == 0 && a[i][j-1] == 0 && a[i-1][j-1] == 0 && a[i-1][j+1] == 0 {
						a[i][j] = 1
					}
				}
			} else {
				for j := 1; j <= w; j += w - 1 {
					if a[i-1][j] == 0 && a[i][j-1] == 0 && a[i-1][j-1] == 0 && a[i-1][j+1] == 0 {
						a[i][j] = 1
					}
				}
			}
		}

		for i := 1; i <= h; i++ {
			for j := 1; j <= w; j++ {
				fmt.Fprint(out, a[i][j])
			}
			fmt.Fprintln(out)
		}
		fmt.Fprintln(out)
	}
}

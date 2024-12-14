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
	a := make([][]int, n)
	for i := 0; i < n; i++ {
		a[i] = make([]int, n)
	}

	for i := 0; i < n/2; i++ {
		d := n/2 - i
		for j := 0; j < d; j++ {
			a[i][j], a[i][n-j-1] = 1, 1
			a[n-i-1][j], a[n-i-1][n-j-1] = 1, 1
		}
	}

	odd, even := 1, 2
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if a[i][j] == 1 {
				a[i][j] = even
				even += 2
			} else {
				a[i][j] = odd
				odd += 2
			}
			fmt.Fprint(out, a[i][j], " ")
		}
		fmt.Fprintln(out)
	}

}

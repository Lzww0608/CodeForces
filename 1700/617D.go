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

	x := make([]int, 3)
	y := make([]int, 3)
	for i := 0; i < 3; i++ {
		fmt.Fscan(in, &x[i], &y[i])
	}

	if x[0] == x[1] && x[1] == x[2] || y[0] == y[1] && y[1] == y[2] {
		fmt.Fprintln(out, 1)
		return
	}

	ans := 3
	for i := 0; i < 3; i++ {
		j, k := (i+1)%3, (i+2)%3
		if x[i] == x[j] {
			if y[k] >= y[i] && y[k] >= y[j] || y[k] <= y[i] && y[k] <= y[j] {
				ans = 2
				break
			}
		} else if y[i] == y[j] {
			if x[k] >= x[i] && x[k] >= x[j] || x[k] <= x[i] && x[k] <= x[j] {
				ans = 2
				break
			}
		}
	}

	fmt.Fprintln(out, ans)
	return
}

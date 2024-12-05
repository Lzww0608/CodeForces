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

	var n, m, k int
	fmt.Fscan(in, &n, &m, &k)
	i, j := 1, 0
	kk := k
	for d := 1; k > 0; k-- {
		if k != 1 {
			fmt.Fprint(out, 2, " ")
			for t := 0; t < 2; t++ {
				j += d
				fmt.Fprint(out, i, j, " ")
				if j == 1 && d == -1 {
					j--
					d = -d
					i++
				} else if j == m && d == 1 {
					j++
					d = -d
					i++
				}
			}
		} else {
			fmt.Fprint(out, n*m-2*(kk-1), " ")
			for t := 0; t < n*m-2*(kk-1); t++ {
				j += d
				fmt.Fprint(out, i, j, " ")
				if j == 1 && d == -1 {
					j--
					d = -d
					i++
				} else if j == m && d == 1 {
					j++
					d = -d
					i++
				}
			}
		}
		fmt.Fprintln(out)
	}
}

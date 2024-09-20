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
	if n <= 2 {
		fmt.Fprintln(out, 1)
		fmt.Fprintln(out, 1)
	} else if n == 3 {
		fmt.Fprintln(out, 2)
		fmt.Fprintln(out, 1, 3)
	} else if n == 4 {
		fmt.Fprintln(out, 4)
		fmt.Fprintln(out, 3, 1, 4, 2)
	} else {
		fmt.Fprintln(out, n)
		for i := 1; i <= (n+1)/2; i++ {
			fmt.Fprint(out, i, " ")
			if i+(n+1)/2 <= n {
				fmt.Fprint(out, i+(n+1)/2, " ")
			}
		}
		fmt.Fprintln(out)
	}
}

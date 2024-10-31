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
	k := 1
	a := make([]int, n+3)
	for i := 2; i <= n; i++ {
		if a[i] == 0 {
			for j := i; j <= n; j += i {
				a[j] = k
			}
			k++
		}
	}
	for _, x := range a[2 : n+1] {
		fmt.Fprint(out, x, " ")
	}
	fmt.Fprintln(out)
}

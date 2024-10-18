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
	sum := 0
	f := true
	for i := range a {
		fmt.Fscan(in, &a[i])
		sum += a[i]
		if a[i] != 0 {
			f = false
		}
	}
	if f {
		fmt.Fprintln(out, "NO")
		return
	}
	fmt.Fprintln(out, "YES")
	if sum != 0 {
		fmt.Fprintln(out, 1)
		fmt.Fprintln(out, 1, n)
		return
	}
	fmt.Fprintln(out, 2)
	for i := range a {
		sum -= a[i]
		if a[i] != 0 && sum != 0 {
			fmt.Fprintln(out, 1, i+1)
			fmt.Fprintln(out, i+2, n)
			break
		}
	}
}

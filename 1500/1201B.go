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
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
		sum += a[i]
	}
	if sum&1 == 1 {
		fmt.Fprintln(out, "NO")
		return
	}

	for _, x := range a {
		if x > sum/2 {
			fmt.Fprintln(out, "NO")
			return
		}
	}
	fmt.Fprintln(out, "YES")
}

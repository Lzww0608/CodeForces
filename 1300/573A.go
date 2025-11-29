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
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
		for a[i]%2 == 0 {
			a[i] /= 2
		}
		for a[i]%3 == 0 {
			a[i] /= 3
		}
	}

	for _, x := range a {
		if x != a[0] {
			fmt.Fprintln(out, "No")
			return
		}
	}
	fmt.Fprintln(out, "Yes")
}

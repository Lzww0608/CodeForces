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

	var t int
	for fmt.Fscan(in, &t); t > 0; t-- {
		solve(in, out)
	}
}

func solve(in *bufio.Reader, out *bufio.Writer) {
	var n int
	fmt.Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		fmt.Fscan(in, &a[i])
	}

	for i := range a {
		for j := 0; j < 5; j++ {
			a[i] += a[i] % 10
		}

		for a[i]%10 != 2 && a[i]%10 != 0 {
			a[i] += a[i] % 10
		}

		if a[i]%10 == 2 {
			a[i] %= 20
		}
	}

	for i := 1; i < n; i++ {
		if a[i] != a[0] {
			fmt.Fprintln(out, "NO")
			return
		}
	}

	fmt.Fprintln(out, "YES")
}

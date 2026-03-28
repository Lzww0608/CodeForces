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

	ans := 0

	for i := 0; i < n; i++ {
		ans++
		j := i + 1
		for j < n && a[j] >= a[j-1] {
			j++
		}
		for j < n && a[j] <= a[j-1] {
			j++
		}
		i = j - 1
	}

	fmt.Fprintln(out, ans)
}

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
	var s string
	fmt.Fscan(in, &s)
	a := []int{}
	for i, j := 0, 0; i < n; i = j {
		for j = i; j < n && s[j] == s[i]; j++ {
		}
		a = append(a, j-i)
	}

	ans := 0
	n = len(a)
	for i, j := 0, 0; i < n; i++ {
		if j < i {
			j = i
		}

		for j < n && a[j] == 1 {
			j++
		}

		if j < n && a[j] > 1 {
			a[j] -= 1
		} else {
			i++
		}

		ans++
	}

	fmt.Fprintln(out, ans)
}

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
	p := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &p[i])
	}
	var s string
	fmt.Fscan(in, &s)
	sum := 0
	for i := range s {
		if s[i] == 'B' {
			sum += p[i]
		}
	}
	ans := sum
	t := sum
	for i := n - 1; i >= 0; i-- {
		if s[i] == 'B' {
			sum -= p[i]
		} else {
			sum += p[i]
		}
		ans = max(ans, sum)
	}
	sum = t
	for i := 0; i < n; i++ {
		if s[i] == 'B' {
			sum -= p[i]
		} else {
			sum += p[i]
		}
		ans = max(ans, sum)
	}

	fmt.Fprintln(out, ans)

}

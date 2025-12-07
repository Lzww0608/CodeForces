package main

import (
	"bufio"
	"fmt"
	"os"
)

const N int = 26

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, m int
	fmt.Fscan(in, &n, &m)
	a := make([]int, n)
	for i := range a {
		fmt.Fscan(in, &a[i])
	}

	if n > m {
		fmt.Fprintln(out, 0)
	} else {
		ans := 1
		for i := range n {
			for j := i + 1; j < n; j++ {
				ans = ans * abs(a[i]-a[j]) % m
			}
		}
		fmt.Fprintln(out, ans)
	}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

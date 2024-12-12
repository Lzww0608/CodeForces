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

	var s string
	fmt.Fscan(in, &s)
	n := len(s)
	ans := make([]int, n)
	for i := 1; i < n; i++ {
		if s[i] == 'a' {
			ans[i] ^= 1
			ans[i-1] ^= 1
		}
	}
	for _, x := range ans {
		fmt.Fprintln(out, x)
	}
}

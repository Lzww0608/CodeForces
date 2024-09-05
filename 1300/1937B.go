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

	var t, n int
	for fmt.Fscan(in, &t); t > 0; t-- {
		fmt.Fscan(in, &n)
		var a, b string
		fmt.Fscan(in, &a)
		fmt.Fscan(in, &b)
		i := 1
		for i = 1; i < n; i++ {
			if a[i] > b[i-1] {
				break
			}
		}
		ans := 1
		for j := i - 1; j-1 >= 0 && a[j] == b[j-1]; j-- {
			ans++
		}
		fmt.Fprintln(out, a[:i]+b[i-1:])
		fmt.Fprintln(out, ans)
	}
}

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

	var n, k int
	fmt.Fscan(in, &n, &k)
	if n == k {
		fmt.Fprintln(out, -1)
		return
	}
	if k == n-1 {
		for i := 1; i <= n; i++ {
			fmt.Print(i, " ")
		}
		fmt.Println()
	} else {
		fmt.Print(n-k, " ")
		for i := 1; i < n-k; i++ {
			fmt.Print(i, " ")
		}
		for i := n - k + 1; i <= n; i++ {
			fmt.Print(i, " ")
		}
		fmt.Println()
	}
}

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
	ans := 1
	// 2n - 2 n - 1
	for i := 0; i < n-1; i++ {
		ans *= 2*n - 2 - i
	}
	for i := 0; i < n-1; i++ {
		ans /= i + 1
	}
	fmt.Fprintln(out, ans)
}

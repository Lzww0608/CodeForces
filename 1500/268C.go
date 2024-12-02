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

	var n, m int
	fmt.Fscan(in, &n, &m)
	ans := min(n+1, m+1)
	fmt.Fprintln(out, ans)
	for i := 0; i < ans; i++ {
		fmt.Fprintln(out, i, ans-i-1)
	}
}

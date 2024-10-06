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
	m := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &m[i])
	}
	vis := map[string]bool{}
	for i := n - 1; i >= 0; i-- {
		if !vis[m[i]] {
			vis[m[i]] = true
			fmt.Fprintln(out, m[i])
		}
	}

}

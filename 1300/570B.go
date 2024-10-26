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
	if n == 1 {
		fmt.Fprintln(out, 1)
		return
	}
	if m-1 >= n-m {
		fmt.Fprintln(out, m-1)
	} else {
		fmt.Fprintln(out, m+1)
	}
}

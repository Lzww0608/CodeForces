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
	mn := max(0, n-m*2)
	mx := 0
	for mx <= n && mx*(mx-1) < m*2 {
		mx++
	}
	fmt.Fprintln(out, mn, n-mx)

}

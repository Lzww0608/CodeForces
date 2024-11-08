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
	if n > m+1 || (n+1)*2 < m {
		fmt.Fprintln(out, -1)
		return
	}
	for n > 0 {
		n--
		for j := 0; j < 2 && m > n; j++ {
			m--
			fmt.Fprint(out, 1)
		}
		fmt.Fprint(out, 0)
	}
	for m > 0 {
		m--
		fmt.Fprint(out, 1)
	}
}

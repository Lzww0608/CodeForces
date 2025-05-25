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

	ans := 0
	for m > n {
		if m&1 == 0 {
			m /= 2
		} else {
			m++
		}
		ans++
	}
	fmt.Fprintln(out, ans+n-m)
}

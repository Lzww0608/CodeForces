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
	fmt.Fprintln(out, solve(n))
}

func solve(n int) (ans int) {
	for x := 1; n > 0; x <<= 1 {
		ans++
		n -= x
	}

	return
}

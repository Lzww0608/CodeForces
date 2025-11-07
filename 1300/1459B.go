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

	if n&1 == 0 {
		fmt.Fprintln(out, (n/2+1)*(n/2+1))
	} else {
		fmt.Fprintln(out, 2*(n/2+1)*(n/2+2))
	}
}

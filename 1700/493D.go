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
		fmt.Fprintln(out, "white")
		fmt.Fprintln(out, 1, 2)
	} else {
		fmt.Fprintln(out, "black")
	}
}

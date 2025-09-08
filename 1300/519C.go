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

	var a, b int
	fmt.Fscan(in, &a, &b)

	if a > b {
		a, b = b, a
	}

	if a*2 <= b {
		fmt.Fprintln(out, a)
	} else {
		fmt.Fprintln(out, (a+b)/3)
	}
}

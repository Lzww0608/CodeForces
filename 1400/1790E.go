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

	var t int
	for fmt.Fscan(in, &t); t > 0; t-- {
		solve(in, out)
	}
}

func solve(in *bufio.Reader, out *bufio.Writer) {
	var x int
	fmt.Fscan(in, &x)
	a, b := 3*x/2, x/2
	if (a+b)/2 == a^b {
		fmt.Fprintln(out, a, b)
	} else {
		fmt.Fprintln(out, -1)
	}
}

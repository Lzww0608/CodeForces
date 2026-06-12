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
	var n int
	fmt.Fscan(in, &n)
	a, b := 0, 0
	var x int
	for range n {
		fmt.Fscan(in, &x)
		if x&1 == 0 {
			a++
		} else {
			b++
		}
	}

	if b%4 == 2 {
		fmt.Fprintln(out, "Bob")
	} else if b%4 == 3 {
		fmt.Fprintln(out, "Alice")
	} else if b%4 == 0 {
		fmt.Fprintln(out, "Alice")
	} else {
		if a&1 == 1 {
			fmt.Fprintln(out, "Alice")
		} else {
			fmt.Fprintln(out, "Bob")
		}
	}
}

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

	var s string
	fmt.Fscanln(in, &s)
	n := len(s)
	fmt.Fprintln(out, 3)
	fmt.Fprintln(out, "L", 2)
	fmt.Fprintln(out, "R", 2)
	fmt.Fprintln(out, "R", n*2-1)
}

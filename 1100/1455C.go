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

	var t, x, y int
	for fmt.Fscan(in, &t); t > 0; t-- {
		fmt.Fscan(in, &x, &y)
		fmt.Fprintln(out, x-1, y)
	}
}

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
	fmt.Fprintln(out, 1)
	pos := 1
	for i := 1; i < n; i++ {
		if i&1 == 1 {
			pos += n - i
			fmt.Fprintln(out, pos)
		} else {
			pos -= n - i
			fmt.Fprintln(out, pos)
		}
	}
}

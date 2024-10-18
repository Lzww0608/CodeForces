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

	var k int
	fmt.Fscan(in, &k)
	if k > 36 {
		fmt.Fprintln(out, -1)
		return
	}
	for k > 0 {
		if k > 1 {
			fmt.Fprint(out, 8)
		} else {
			fmt.Fprint(out, 6)
		}
		k -= 2
	}
	fmt.Fprintln(out)
}

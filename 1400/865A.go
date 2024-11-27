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

	var a int
	fmt.Fscan(in, &a)
	if a == 1 {
		fmt.Fprintln(out, 1, 1)
		fmt.Fprintln(out, 1)
		return
	}
	fmt.Fprintln(out, (a-1)*2, 2)
	fmt.Fprintln(out, 1, 2)
}

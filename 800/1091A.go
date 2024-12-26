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

	var a, b, c int
	fmt.Fscan(in, &a, &b, &c)
	ans := min(a, b-1, c-2)
	fmt.Fprintln(out, ans*3+3)
}

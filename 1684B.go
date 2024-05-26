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
	var t, a, b, c int64
	for fmt.Fscan(in, &t); t > 0; t-- {
		fmt.Fscan(in, &a, &b, &c)
		fmt.Fprintln(out, a+b*c, b, c)
	}
}

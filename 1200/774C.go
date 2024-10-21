package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)
	if n&1 == 1 {
		n -= 3
		fmt.Fprint(out, "7")
	}
	if n > 0 {
		fmt.Fprintln(out, strings.Repeat("1", n/2))
	}

}

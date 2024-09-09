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
	var n1, n2, k1, k2 int
	fmt.Fscan(in, &n1, &n2, &k1, &k2)
	if n1 <= n2 {
		fmt.Fprintln(out, "Second")
	} else {
		fmt.Fprintln(out, "First")
	}

}

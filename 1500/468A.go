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
	if n < 4 {
		fmt.Fprintln(out, "NO")
		return
	}
	fmt.Fprintln(out, "YES")
	for i := n; i > 5; i -= 2 {
		fmt.Fprintf(out, "%d - %d = %d\n", i, i-1, 1)
		fmt.Fprintf(out, "%d * %d = %d\n", 1, 1, 1)
	}

	if n&1 == 1 {
		fmt.Fprintf(out, "%d - %d = %d\n", 5, 3, 2)
		fmt.Fprintf(out, "%d * %d = %d\n", 2, 4, 8)
		fmt.Fprintf(out, "%d + %d = %d\n", 1, 2, 3)
		fmt.Fprintf(out, "%d * %d = %d\n", 3, 8, 24)
	} else {
		fmt.Fprintf(out, "%d * %d = %d\n", 1, 2, 2)
		fmt.Fprintf(out, "%d * %d = %d\n", 2, 3, 6)
		fmt.Fprintf(out, "%d * %d = %d\n", 6, 4, 24)
	}

}

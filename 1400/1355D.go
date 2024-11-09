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

	var n, s int
	fmt.Fscan(in, &n, &s)
	if s >= n*2 {
		fmt.Fprintln(out, "YES")
		for i := 0; i < n-1; i++ {
			fmt.Fprint(out, 2, " ")
		}
		fmt.Fprintln(out, s-(n-1)*2)
		fmt.Fprintln(out, 1)
	} else {
		fmt.Fprintln(out, "NO")
	}
}

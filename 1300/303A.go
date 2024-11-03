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
	if n&1 == 0 {
		fmt.Fprintln(out, -1)
		return
	}
	for i := 0; i < n; i++ {
		fmt.Fprint(out, i, " ")
	}
	fmt.Fprintln(out)
	for i := 1; i <= n; i++ {
		fmt.Fprint(out, i%n, " ")
	}
	fmt.Fprintln(out)
	for i := 0; i < n; i++ {
		fmt.Fprint(out, (i+(i+1)%n)%n, " ")
	}

}

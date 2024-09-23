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

	var n, x int
	fmt.Fscan(in, &n)
	extra := 0
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &x)
		if x&1 == 1 {
			extra ^= 1
		} else if x == 0 {
			if extra&1 == 1 {
				fmt.Fprintln(out, "NO")
				return
			}
		}
	}
	if extra == 0 {
		fmt.Fprintln(out, "YES")
	} else {
		fmt.Fprintln(out, "NO")
	}
}

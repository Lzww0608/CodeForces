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

	var t, l, r int64
	for fmt.Fscan(in, &t); t > 0; t-- {
		fmt.Fscan(in, &l, &r)
		if l*2 >= r+1 {
			fmt.Fprintln(out, "YES")
		} else {
			fmt.Fprintln(out, "NO")
		}
	}
}

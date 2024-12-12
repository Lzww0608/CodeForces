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
	var s string
	fmt.Fscan(in, &s)
	cnt := 0
	for i := 0; i < n; i++ {
		if s[i] == ')' {
			cnt--
		}
		if cnt&1 == 0 {
			fmt.Fprint(out, 1)
		} else {
			fmt.Fprint(out, 0)
		}
		if s[i] == '(' {
			cnt++
		}
	}
}

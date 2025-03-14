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

	var a, b string
	fmt.Fscan(in, &a, &b)
	cntA, cntB := 0, 0
	for i := range a {
		if a[i] == '1' {
			cntA++
		}
	}
	for i := range b {
		if b[i] == '1' {
			cntB++
		}
	}
	if cntA+cntA&1 >= cntB {
		fmt.Fprintln(out, "YES")
	} else {
		fmt.Fprintln(out, "NO")
	}
}

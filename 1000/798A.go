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

	var s string
	fmt.Fscan(in, &s)
	n := len(s)
	cnt := 0
	for i := 0; i+i+1 < n; i++ {
		if s[i] != s[n-i-1] {
			cnt++
		}
	}
	if cnt == 1 || (cnt == 0 && n&1 == 1) {
		fmt.Fprintln(out, "YES")
	} else {
		fmt.Fprintln(out, "NO")
	}
}

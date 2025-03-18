package main

import (
	"bufio"
	"fmt"
	"os"
)

const MOD int = 998244353

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)
	var s string
	fmt.Fscan(in, &s)

	l, r := 1, 1
	for i := 0; i < n && s[i] == s[0]; i++ {
		l++
	}
	for i := n - 1; i >= 0 && s[i] == s[n-1]; i-- {
		r++
	}

	if s[0] == s[n-1] {
		fmt.Fprintln(out, l*r%MOD)
	} else {
		fmt.Fprintln(out, (l+r-1)%MOD)
	}

}

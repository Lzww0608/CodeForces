package main

import (
	"bufio"
	"fmt"
	"os"
)

const MOD int = 1_000_000_007

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var s string
	fmt.Fscan(in, &s)
	n := len(s)

	f := make([]int, n+1)
	f[0] = 1
	for i := range s {
		if s[i] == 'w' || s[i] == 'm' {
			fmt.Fprintln(out, 0)
			return
		}
		if i > 0 && (s[i] == 'u' && s[i-1] == 'u' || s[i] == 'n' && s[i-1] == 'n') {
			f[i+1] = (f[i-1] + f[i]) % MOD
		} else {
			f[i+1] = f[i]
		}
	}

	fmt.Fprintln(out, f[n])
}

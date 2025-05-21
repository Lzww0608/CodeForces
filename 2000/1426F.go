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

	var n int
	fmt.Fscan(in, &n)
	var s string
	fmt.Fscan(in, &s)
	f := [4]int{}
	f[0] = 1
	for _, c := range s {
		if c == 'c' {
			f[3] = (f[3] + f[2]) % MOD
		} else if c == 'b' {
			f[2] = (f[2] + f[1]) % MOD
		} else if c == 'a' {
			f[1] = (f[1] + f[0]) % MOD
		} else {
			for i := 3; i > 0; i-- {
				f[i] = (3*f[i] + f[i-1]) % MOD
			}
			f[0] = 3 * f[0] % MOD
		}
	}

	fmt.Fprintln(out, f[3])
}

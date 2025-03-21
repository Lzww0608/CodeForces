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

	var n, l, r int
	fmt.Fscan(in, &n, &l, &r)
	f := make([][3]int, n)
	rem := [3]int{(r - l + 1) / 3, (r - l + 1) / 3, (r - l + 1) / 3}
	if (r-l+1)%3 == 1 {
		rem[r%3]++
	} else if (r-l+1)%3 == 2 {
		rem[l%3]++
		rem[(l+1)%3]++
	}

	f[0] = rem
	for i := 1; i < n; i++ {
		f[i][0] = (f[i-1][0]*rem[0] + f[i-1][1]*rem[2] + f[i-1][2]*rem[1]) % MOD
		f[i][1] = (f[i-1][0]*rem[1] + f[i-1][1]*rem[0] + f[i-1][2]*rem[2]) % MOD
		f[i][2] = (f[i-1][0]*rem[2] + f[i-1][1]*rem[1] + f[i-1][2]*rem[0]) % MOD
	}

	fmt.Fprintln(out, f[n-1][0]%MOD)
}

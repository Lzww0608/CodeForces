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
	a, b, c, d := 1, 0, 0, 0
	for i := 1; i <= n; i++ {
		a, b, c, d = (b+c+d)%MOD, (a+c+d)%MOD, (a+b+d)%MOD, (a+b+c)%MOD
	}

	fmt.Fprintln(out, a)
}

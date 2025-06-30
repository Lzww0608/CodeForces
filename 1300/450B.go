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

	var x, y, n, ans int
	fmt.Fscan(in, &x, &y, &n)
	n--
	switch n % 6 {
	case 0:
		ans = x + MOD
	case 1:
		ans = y + MOD
	case 2:
		ans = y - x + MOD
	case 3:
		ans = -x + MOD
	case 4:
		ans = -y + MOD
	case 5:
		ans = x - y + MOD
	}
	fmt.Fprintln(out, (ans+MOD)%MOD)
}

package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, x int
	fmt.Fscan(in, &n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &x)
		ans := 15
		for j := 0; j < 15; j++ {
			ans = min(ans, j+15-min(15, bits.TrailingZeros(uint(x+j))))
		}

		fmt.Fprint(out, ans, " ")
	}
}

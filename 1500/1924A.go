package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"
	"strings"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var t, n, k, m int

	for fmt.Fscan(in, &t); t > 0; t-- {
		fmt.Fscan(in, &n, &k, &m)
		var s string
		fmt.Fscan(in, &s)
		mask := 0
		d := 1
		a := []rune{}
		for _, c := range s {
			x := int(c - 'a')
			mask |= 1 << x
			if mask == 1<<k-1 {
				mask = 0
				a = append(a, c)
				d++
			}
		}

		if d > n {
			fmt.Fprintln(out, "YES")
		} else {
			fmt.Fprintln(out, "NO")
			idx := 0
			for (mask>>idx)&1 == 1 {
				idx++
			}
			fmt.Fprintln(out, string(a)+string('a'+bits.TrailingZeros(uint(^mask)))+strings.Repeat("a", n-1-len(a)))
		}
	}
}

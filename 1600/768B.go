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

	var n, l, r uint64
	fmt.Fscan(in, &n, &l, &r)
	mx := ((uint64(1) << bits.Len64(n)) - 1) >> 1

	f := func(x uint64) (res uint64) {
		for i, cur := n, mx; x > 0; i >>= 1 {
			if x >= cur {
				x -= cur
				res += i / 2
				if x > 0 {
					res += i & 1
					x--
				}
			}
			cur >>= 1
		}
		return
	}

	ans := f(r) - f(l-1)
	fmt.Fprintln(out, ans)
}

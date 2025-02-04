package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"
)

const N int = 32

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var t int
	for fmt.Fscan(in, &t); t > 0; t-- {
		solve(in, out)
	}
}

func solve(in *bufio.Reader, out *bufio.Writer) {
	var n int
	fmt.Fscan(in, &n)
	a := make([]int, n)
	sum := 0
	for i := range a {
		fmt.Fscan(in, &a[i])
		sum += a[i]
	}

	if sum%n != 0 {
		fmt.Fprintln(out, "No")
		return
	}
	sum /= n
	cnt := make([]int, N)
	for _, x := range a {
		if x == sum {
			continue
		}
		d := abs(x - sum)
		l := bits.TrailingZeros(uint(d))
		if x > sum {
			cnt[l]++
		} else {
			cnt[l]--
		}
		d += 1 << l
		if d&(d-1) != 0 {
			fmt.Fprintln(out, "No")
			return
		}
		l = bits.TrailingZeros(uint(d))
		if x > sum {
			cnt[l]--
		} else {
			cnt[l]++
		}
	}
	for _, x := range cnt {
		if x != 0 {
			fmt.Fprintln(out, "No")
			return
		}
	}
	fmt.Fprintln(out, "Yes")
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

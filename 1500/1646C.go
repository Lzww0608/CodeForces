package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"
)

var cnt []int

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	for i, x := 3, 2; x <= int(1e12); i++ {
		x *= i
		cnt = append(cnt, x)
	}

	var t int
	for fmt.Fscan(in, &t); t > 0; t-- {
		solve(in, out)
	}

}

func solve(in *bufio.Reader, out *bufio.Writer) {
	var n int
	fmt.Fscan(in, &n)

	ans := bits.OnesCount(uint(n))
	m := len(cnt)
	for i := 1; i < (1 << m); i++ {
		x := bits.OnesCount(uint(i))
		t := n
		for j, y := range cnt {
			if i&(1<<j) != 0 {
				t -= y
			}
		}
		ans = min(ans, x+bits.OnesCount(uint(t)))
	}

	fmt.Fprintln(out, ans)
}

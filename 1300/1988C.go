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

	var t int
	for fmt.Fscan(in, &t); t > 0; t-- {
		solve(in, out)
	}

}

func solve(in *bufio.Reader, out *bufio.Writer) {
	var n int
	fmt.Fscan(in, &n)
	k := bits.OnesCount(uint(n)) + 1
	if k == 2 {
		fmt.Fprintln(out, 1)
		fmt.Fprintln(out, n)
		return
	}
	fmt.Fprintln(out, k)
	t := n
	ans := make([]int, 0, k)
	for i := 1; i < k; i++ {
		ans = append(ans, n-t&(-t))
		t -= t & (-t)
	}
	m := len(ans)
	for i := m - 1; i >= 0; i-- {
		fmt.Fprint(out, ans[i], " ")
	}
	fmt.Fprintln(out, n)
}

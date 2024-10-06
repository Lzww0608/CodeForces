package main

import (
	"bufio"
	"fmt"
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
	var n, k int
	fmt.Fscan(in, &n, &k)
	a := make([]int, n)
	sum, cnt := 0, 0
	for i := range a {
		fmt.Fscan(in, &a[i])
		sum += a[i]
		if sum&1 != cnt&1 {
			cnt++
		}
	}
	if sum&1 != k&1 || cnt < k {
		fmt.Fprintln(out, "NO")
		return
	}
	fmt.Fprintln(out, "YES")
	sum, cnt = 0, 0
	if k == 1 {
		fmt.Fprintln(out, n)
		return
	}
	for i, x := range a {
		sum += x
		if sum&1 != cnt&1 {
			cnt++
			fmt.Fprint(out, i+1, " ")
		}

		if cnt == k-1 {
			break
		}
	}
	fmt.Fprintln(out, n)
}

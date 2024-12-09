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
	var n int
	fmt.Fscan(in, &n)
	var s string
	fmt.Fscan(in, &s)

	ans := n * (n + 1) / 2
	need := 0
	for i := n - 1; i >= 0; i-- {
		if s[i] == '1' && need+1 <= i {
			ans -= i + 1
			need++
		} else {
			need = max(0, need-1)
		}
	}
	fmt.Fprintln(out, ans)
}

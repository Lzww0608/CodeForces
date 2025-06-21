package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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
	ans := []byte(strings.Repeat("a", n))
	for i := 1; ; i++ {
		if k > i {
			k -= i
		} else {
			ans[n-1-i] = 'b'
			ans[n-k] = 'b'
			fmt.Fprintln(out, string(ans))
			break
		}
	}

}

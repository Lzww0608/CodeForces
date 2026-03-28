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
	var k int
	fmt.Fscan(in, &k)

	ans := []byte{}
	for k > 0 {
		x := k % 9
		if x < 4 {
			ans = append(ans, byte(x+'0'))
		} else {
			ans = append(ans, byte(x+'0'+1))
		}
		k /= 9
	}

	for i := len(ans) - 1; i >= 0; i-- {
		fmt.Fprint(out, string(ans[i]))
	}
	fmt.Fprintln(out)
}

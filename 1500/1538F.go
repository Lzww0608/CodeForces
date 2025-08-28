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
	var l, r int
	fmt.Fscan(in, &l, &r)

	ans := 0
	for l > 0 || r > 0 {
		ans += r - l
		r /= 10
		l /= 10
	}

	fmt.Fprintln(out, ans)
}

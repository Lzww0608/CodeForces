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
	var n int
	fmt.Fscan(in, &n)
	if n == 1 {
		fmt.Fprintln(out, 1)
		return
	} else if n == 3 {
		fmt.Fprintln(out, 169)
		fmt.Fprintln(out, 196)
		fmt.Fprintln(out, 961)
		return
	}
	fmt.Fprint(out, 196)
	fmt.Fprintln(out, strings.Repeat("0", n-3))
	for i := 0; i < n/2; i++ {
		fmt.Fprint(out, 1)
		fmt.Fprint(out, strings.Repeat("0", i))
		fmt.Fprint(out, 6)
		fmt.Fprint(out, strings.Repeat("0", i))
		fmt.Fprint(out, 9)
		fmt.Fprintln(out, strings.Repeat("0", n-i*2-3))

		fmt.Fprint(out, 9)
		fmt.Fprint(out, strings.Repeat("0", i))
		fmt.Fprint(out, 6)
		fmt.Fprint(out, strings.Repeat("0", i))
		fmt.Fprint(out, 1)
		fmt.Fprintln(out, strings.Repeat("0", n-i*2-3))
	}
}

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
	var n0, n1, n2 int
	fmt.Fscan(in, &n0, &n1, &n2)
	if n1 == 0 {
		if n0 > 0 {
			fmt.Fprintln(out, strings.Repeat("0", n0+1))
		} else {
			fmt.Fprintln(out, strings.Repeat("1", n2+1))
		}
		return
	}
	pre := -1
	if n0 > 0 {
		fmt.Fprint(out, strings.Repeat("0", n0+1))
		pre = 0
	}

	if n2 > 0 {
		fmt.Fprint(out, strings.Repeat("1", n2+1))
		pre = 1
	}

	if n0 > 0 && n2 > 0 {
		n1--
	}

	if pre == -1 {
		fmt.Fprint(out, "0")
		pre = 0
	}
	for n1 > 0 {
		n1--
		if pre == 0 {
			fmt.Fprint(out, "1")
		} else {
			fmt.Fprint(out, "0")
		}
		pre ^= 1
	}
	fmt.Fprintln(out)
}

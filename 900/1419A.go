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
	fmt.Fscan(in, &t)
	for t > 0 {
		solve(in, out)
		t--
	}
}

func solve(in *bufio.Reader, out *bufio.Writer) {
	var n int
	fmt.Fscan(in, &n)
	var s string
	fmt.Fscan(in, &s)
	if n&1 == 0 {
		even := 0
		for i := 1; i < n; i += 2 {
			x := s[i] - '0'
			if x&1 == 0 {
				even++
			}
		}
		if even > 0 {
			fmt.Fprintln(out, "2")
		} else {
			fmt.Fprintln(out, "1")
		}
	} else {
		odd := 0
		for i := 0; i < n; i += 2 {
			x := s[i] - '0'
			if x&1 == 1 {
				odd++
			}
		}
		if odd > 0 {
			fmt.Fprintln(out, "1")
		} else {
			fmt.Fprintln(out, "2")
		}
	}
}

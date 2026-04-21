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
	if n&1 == 1 {
		fmt.Fprintln(out, "Bob")
		return
	}

	cnt := 0
	for n&1 == 0 {
		cnt++
		n >>= 1
	}

	if n > 1 {
		fmt.Fprintln(out, "Alice")
	} else {
		if cnt%2 == 0 {
			fmt.Fprintln(out, "Alice")
		} else {
			fmt.Fprintln(out, "Bob")
		}
	}
}

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

	check := func(mid int) bool {
		for mid > 0 {
			if mid%3 > 1 {
				return false
			}
			mid /= 3
		}
		return true
	}

	for i := n; ; i++ {
		if check(i) {
			fmt.Fprintln(out, i)
			return
		}
	}

}

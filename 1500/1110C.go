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

	var q int
	for fmt.Fscan(in, &q); q > 0; q-- {
		solve(in, out)
	}
}

func solve(in *bufio.Reader, out *bufio.Writer) {
	var n int
	fmt.Fscan(in, &n)
	x := 1
	for x <= n {
		x <<= 1
	}
	if n == x-1 {
		for b := 2; b*b <= n; b++ {
			if n%b == 0 {
				fmt.Fprintln(out, n/b)
				return
			}
		}
		fmt.Fprintln(out, 1)
	} else {
		fmt.Fprintln(out, x-1)
	}

}

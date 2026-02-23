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

	ans := 1
	for n > 0 {
		t := n % 10
		n = n / 10

		cur := 0
		for i := range t + 1 {
			for j := range t + 1 {
				k := t - i - j
				if k >= 0 {
					cur++
				}
			}
		}
		ans *= cur
	}
	fmt.Fprintln(out, ans)
}

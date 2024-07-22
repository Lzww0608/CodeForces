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

	var t, d int
	for fmt.Fscan(in, &t); t > 0; t-- {
		fmt.Fscan(in, &d)
		if d == 1 {
			fmt.Fprintln(out, 6)
		} else if d == 2 {
			fmt.Fprintln(out, 15)
		} else {
			ans := 1
			x := 1 + d
			for !check(x) {
				x++
			}
			ans *= x
			x += d
			for !check(x) {
				x++
			}
			fmt.Fprintln(out, x*ans)
		}
	}
}

func check(x int) bool {
	for i := 2; i*i <= x; i++ {
		if x%i == 0 {
			return false
		}
	}
	return true
}

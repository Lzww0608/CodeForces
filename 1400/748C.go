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

	var n int
	fmt.Fscan(in, &n)
	var s string
	fmt.Fscan(in, &s)
	ans := 1
	x, y := 0, 0
	for i := range s {
		if s[i] == 'R' {
			if x < 0 {
				ans++
				x, y = 1, 0
			} else {
				x++
			}
		} else if s[i] == 'L' {
			if x > 0 {
				ans++
				x, y = -1, 0
			} else {
				x--
			}
		} else if s[i] == 'U' {
			if y < 0 {
				ans++
				x, y = 0, 1
			} else {
				y++
			}
		} else {
			if y > 0 {
				ans++
				x, y = 0, -1
			} else {
				y--
			}
		}
	}

	fmt.Fprintln(out, ans)
}

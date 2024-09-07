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

	var t, n int
	for fmt.Fscan(in, &t); t > 0; t-- {
		fmt.Fscan(in, &n)
		var s string
		fmt.Fscan(in, &s)
		ans := 0
		st := []int{}
		for i, c := range s {
			if c == '_' {
				if len(st) > 0 {
					ans += i - st[len(st)-1]
					st = st[:len(st)-1]
				} else {
					st = append(st, i)
				}
			} else if c == '(' {
				st = append(st, i)
			} else {
				ans += i - st[len(st)-1]
				st = st[:len(st)-1]
			}
		}

		fmt.Fprintln(out, ans)
	}
}

package main

import (
	"bufio"
	"fmt"
	"os"
)

// https://codeforces.com/problemset/problem/343/B
func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var s string
	fmt.Fscan(in, &s)
	st := []byte{}
	for i := range s {
		if len(st) > 0 && st[len(st)-1] == s[i] {
			st = st[:len(st)-1]
		} else {
			st = append(st, s[i])
		}
	}

	if len(st) == 0 {
		fmt.Fprintln(out, "Yes")
	} else {
		fmt.Fprintln(out, "No")
	}
}

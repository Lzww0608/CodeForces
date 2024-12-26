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

	var s, t string
	fmt.Fscan(in, &s, &t)
	n := len(s)
	ans1 := make([]byte, n)
	ans2 := make([]byte, n)
	f1, f2 := true, true
	for i := 0; i < n; i++ {
		if f1 {
			if s[i] == t[i] {
				ans1[i] = s[i]
			} else {
				ans1[i] = s[i] + 1
				f1 = false
			}
		} else {
			ans1[i] = 'a'
		}

		if f2 {
			ans2[i] = s[i]
			if s[i] != t[i] {
				f2 = false
			}
		} else {
			ans2[i] = 'z'
		}
	}
	a := string(ans1)
	b := string(ans2)
	if a > s && a < t {
		fmt.Fprintln(out, a)
	} else if b > s && b < t {
		fmt.Fprintln(out, b)
	} else {
		fmt.Fprintln(out, "No such string")
	}
}

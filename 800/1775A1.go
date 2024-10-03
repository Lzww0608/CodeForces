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
next:
	for fmt.Fscan(in, &t); t > 0; t-- {
		var s string
		fmt.Fscan(in, &s)
		n := len(s)
		for i := 0; i < n-2; i++ {
			for j := i + 1; j < n-1; j++ {
				if s[:i+1] <= s[i+1:j+1] && s[j+1:] <= s[i+1:j+1] || s[:i+1] >= s[i+1:j+1] && s[j+1:] >= s[i+1:j+1] {
					fmt.Fprint(out, s[:i+1], " ")
					fmt.Fprint(out, s[i+1:j+1], " ")
					fmt.Fprintln(out, s[j+1:])
					continue next
				}
			}
		}
		fmt.Fprintln(out, ":(")
	}

}

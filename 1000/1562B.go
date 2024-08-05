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

	var t, k int
	m := map[int]bool{
		1: true,
		4: true,
		6: true,
		8: true,
		9: true,
	}
next:
	for fmt.Fscan(in, &t); t > 0; t-- {
		fmt.Fscan(in, &k)
		var s string
		fmt.Fscan(in, &s)
		for i := range s {
			if m[int(s[i]-'0')] {
				fmt.Fprintln(out, 1)
				fmt.Fprintln(out, int(s[i]-'0'))
				continue next
			}
		}
		if s[:2] != "53" && s[:2] != "23" && s[:2] != "73" && s[:2] != "37" {
			fmt.Fprintln(out, 2)
			fmt.Fprintln(out, int(s[0]-'0')*10+int(s[1]-'0'))
			continue
		} else if s[1:3] != "53" && s[1:3] != "23" && s[1:3] != "73" && s[1:3] != "37" {
			fmt.Fprintln(out, 2)
			fmt.Fprintln(out, int(s[1]-'0')*10+int(s[2]-'0'))
			continue
		} else {
			fmt.Fprintln(out, 2)
			fmt.Fprintln(out, int(s[0]-'0')*10+int(s[2]-'0'))
		}

	}
}

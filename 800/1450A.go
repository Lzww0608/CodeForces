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
		a := [26]int{}
		for _, c := range s {
			a[int(c-'a')]++
		}
		for i, x := range a {
			for j := 0; j < x; j++ {
				fmt.Fprint(out, string(i+'a'))
			}

		}
		fmt.Fprintln(out)
	}

}

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
		var s string
		fmt.Fscan(in, &s)
		cnt := [26]int{}
		n := 0
		for _, c := range s {
			cnt[int(c-'a')]++
			if cnt[int(c-'a')] == 1 {
				n++
			}
		}
		if n == 1 {
			fmt.Fprintln(out, -1)
		} else {
			for i, x := range cnt {
				for x > 0 {
					x--
					fmt.Fprint(out, string('a'+i))
				}
			}
			fmt.Fprintln(out)
		}

	}
}

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
		a := make([]int, n)
		for i := range a {
			fmt.Fscan(in, &a[i])
		}

		cnt := [26]int{}
		for _, x := range a {
			for i, y := range cnt {
				if y == x {
					cnt[i]++
					fmt.Fprint(out, string('a'+i))
					break
				}
			}
		}
		fmt.Fprintln(out)

	}

}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var k int
	fmt.Fscan(in, &k)
	for i := 5; i*i <= k; i++ {
		if k%i == 0 && k/i >= 5 {
			j := k / i
			for t := 0; t < i; t++ {
				if t%5 == 0 {
					fmt.Fprint(out, "aeiou")
					fmt.Fprint(out, strings.Repeat("a", j-5))
				} else if t%5 == 1 {
					fmt.Fprint(out, "eioua")
					fmt.Fprint(out, strings.Repeat("e", j-5))
				} else if t%5 == 2 {
					fmt.Fprint(out, "iouae")
					fmt.Fprint(out, strings.Repeat("i", j-5))
				} else if t%5 == 3 {
					fmt.Fprint(out, "ouaei")
					fmt.Fprint(out, strings.Repeat("o", j-5))
				} else {
					fmt.Fprint(out, "uaeio")
					fmt.Fprint(out, strings.Repeat("u", j-5))
				}
			}
			fmt.Fprintln(out)
			return
		}
	}
	fmt.Fprintln(out, -1)
}

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
		var s, t string
		fmt.Fscan(in, &s)
		fmt.Fscan(in, &t)
		freq := make([]int, 26)
		for _, c := range s {
			freq[int(c-'a')]++
		}

		if t == "abc" && freq[0] > 0 {
			for freq[0] > 0 {
				fmt.Fprint(out, "a")
				freq[0]--
			}
			for freq[2] > 0 {
				fmt.Fprint(out, "c")
				freq[2]--
			}
		}

		for i, x := range freq {
			for x > 0 {
				x--
				fmt.Fprint(out, string('a'+i))
			}
		}
		fmt.Fprintln(out)
	}
}

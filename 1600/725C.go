package main

import (
	"bufio"
	"fmt"
	"os"
)

const N int = 26

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var s []byte
	fmt.Fscan(in, &s)
	pos := map[byte]int{}
	for i, c := range s {
		p, ok := pos[c]
		if !ok {
			pos[c] = i
			continue
		}
		d := i - p
		if d == 1 {
			fmt.Fprintln(out, "Impossible")
		} else {
			s = append(append(s[p+d/2+1:], s[:p]...), s[p+1:p+d/2+1]...)
			for i := 0; i < 6; i++ {
				s[i], s[12-i] = s[12-i], s[i]
			}
			fmt.Fprintf(out, "%s\n%s\n", s[:13], s[13:])
		}
		return
	}
}

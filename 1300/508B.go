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

	var str string
	fmt.Fscan(in, &str)
	s := []byte(str)
	n := len(s)

	j := -1
	for i := 0; i < n-1; i++ {
		y := int(s[i] - '0')
		if y&1 == 0 {
			if s[i] < s[n-1] {
				j = i
				break
			} else {
				j = i
			}
		}
	}

	if j == -1 {
		fmt.Fprintln(out, -1)
		return
	}
	s[j], s[n-1] = s[n-1], s[j]
	fmt.Fprintln(out, string(s))

}

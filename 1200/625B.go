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

	ans := 0
	n, m := len(s), len(t)
	for i := 0; i+m <= n; i++ {
		if s[i:i+m] == t {
			ans++
			i = i + m - 1
		}
	}

	fmt.Fprintln(out, ans)
}

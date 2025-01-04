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

	var n int
	fmt.Fscan(in, &n)
	ans := 1
	for k := 1; ; k++ {
		t := ((1 << k) - 1) * (1 << (k - 1))
		if t > n {
			break
		} else if n%t == 0 {
			ans = t
		}
	}

	fmt.Fprintln(out, ans)
}

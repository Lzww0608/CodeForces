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
	a, b := 1, 2
	for a+b <= n {
		ans++
		a, b = b, a+b
	}
	fmt.Fprintln(out, ans)
}

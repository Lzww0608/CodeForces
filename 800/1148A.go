package main

import (
	"bufio"
	"fmt"
	"os"
)

const N int = 26

var id []int

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var a, b, c int
	fmt.Fscan(in, &a, &b, &c)
	d := min(a, b)
	ans := c*2 + d*2
	if a > d || b > d {
		ans++
	}
	fmt.Fprintln(out, ans)
}

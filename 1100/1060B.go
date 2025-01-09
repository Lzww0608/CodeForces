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
	if n < 10 {
		fmt.Fprintln(out, n)
		return
	}
	x := 9
	for x <= n {
		x = x*10 + 9
	}
	x /= 10
	fmt.Fprintln(out, calc(x)+calc(n-x))
}

func calc(x int) (res int) {
	for x > 0 {
		res += x % 10
		x /= 10
	}
	return
}

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
		solve(in, out)
	}
}

func solve(in *bufio.Reader, out *bufio.Writer) {
	var x, y int
	fmt.Fscan(in, &x, &y)
	ans := 0
	if x == y {
		ans = x
	} else if x > y {
		ans = x + y
	} else {
		ans = y - (y%x)/2
	}
	if ans%x != y%ans {
		panic("Wrong!")
	}
	fmt.Fprintln(out, ans)
}

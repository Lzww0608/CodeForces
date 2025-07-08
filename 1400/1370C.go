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

func check(x int) bool {
	for i := 2; i <= x/i; i++ {
		if x%i == 0 {
			return false
		}
	}

	return true
}

func solve(in *bufio.Reader, out *bufio.Writer) {
	var n int
	fmt.Fscan(in, &n)
	f := 0
	if n == 1 {
		f = 1
	} else if n == 2 {
		f = 0
	} else if n&1 == 1 {
		f = 0
	} else if n&(n-1) == 0 {
		f = 1
	} else if n%4 != 0 && check(n/2) {
		f = 1
	}

	if f == 0 {
		fmt.Fprintln(out, "Ashishgup")
	} else {
		fmt.Fprintln(out, "FastestFinger")
	}
}

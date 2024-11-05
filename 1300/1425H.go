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
	ans := make([]bool, 4)
	var a, b, c, d int
	fmt.Fscan(in, &a, &b, &c, &d)
	if (a+b)&1 == 0 {
		if b+c > 0 {
			ans[2] = true
		}
		if a+d > 0 {
			ans[3] = true
		}
	} else {
		if b+c > 0 {
			ans[1] = true
		}
		if a+d > 0 {
			ans[0] = true
		}
	}
	for _, x := range ans {
		if x {
			fmt.Fprint(out, "Ya ")
		} else {
			fmt.Fprint(out, "Tidak ")
		}
	}
	fmt.Fprintln(out)
}

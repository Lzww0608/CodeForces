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
	if n%4 > 1 {
		fmt.Fprintln(out, -1)
		return
	} else if n == 1 {
		fmt.Fprintln(out, 1)
		return
	}

	for i := 1; i <= n/2; i++ {
		if i&1 == 1 {
			fmt.Fprint(out, i+1, " ")
		} else {
			fmt.Fprint(out, n-i+2, " ")
		}
	}

	m := n / 2
	if n%4 == 1 {
		fmt.Fprint(out, m+1, " ")
		m++
	}
	for i := m + 1; i <= n; i++ {
		if (i-m)&1 == 1 {
			fmt.Fprint(out, n-i, " ")
		} else {
			fmt.Fprint(out, i-1, " ")
		}
	}
	//fmt.Fprint(out, n-1, " ")

}

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

	var n, p int
	fmt.Fscan(in, &n)

	if n == 1 {
		fmt.Fprintln(out, 1)
		return
	}

	mn := n * 2
	for i := 1; i <= n; i++ {
		t := i + n/i
		if n%i != 0 {
			t++
		}
		if t < mn {
			mn, p = t, i
		}
	}

	for cur := n - p + 1; cur > 0; cur -= p {
		for i := cur; i < cur+p; i++ {
			fmt.Fprint(out, i, " ")
		}
		if cur > 0 && cur-p <= 0 {
			for i := 1; i < cur; i++ {
				fmt.Fprint(out, i, " ")
			}
		}
	}
	fmt.Fprintln(out)
}

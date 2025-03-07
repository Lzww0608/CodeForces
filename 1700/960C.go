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

	var X, d int
	fmt.Fscan(in, &X, &d)

	ans := solve(X, 1, d)
	fmt.Fprintln(out, len(ans))
	for _, x := range ans {
		fmt.Fprint(out, x, " ")
	}
}

func solve(x, k, d int) (res []int) {
	t := 0
	for (1<<t)-1 < x {
		t++
	}

	if x == (1<<t)-1 {
		for i := 0; i < t; i++ {
			res = append(res, k)
		}
	} else {
		t--
		for i := 0; i < t; i++ {
			res = append(res, k)
		}
		res = append(res, solve(x-(1<<t)+1, k+d, d)...)
	}
	return
}

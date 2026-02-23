package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var v int
	fmt.Fscan(in, &v)
	a := make([]int, 9)
	mn := math.MaxInt32
	for i := range a {
		fmt.Fscan(in, &a[i])
		mn = min(mn, a[i])
	}

	t := v / mn
	if t == 0 {
		fmt.Fprintln(out, -1)
		return
	}

	for ; t >= 0; t-- {
		for i := 9; i > 0; i-- {
			if c := a[i-1]; c <= v && (v-c)/mn == t {
				v -= c
				fmt.Fprint(out, i)
				break
			}
		}
	}
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var a, b, c int
	fmt.Fscan(in, &a, &b, &c)

	ans := []int{}
	for s := 0; s <= 81; s++ {
		x := b*pow(s, a) + c
		y := x
		t := 0
		for x > 0 {
			t += x % 10
			x /= 10
		}
		if t == s && y < 1e9 && y > 0 {
			ans = append(ans, y)
		}
	}

	sort.Ints(ans)
	fmt.Fprintln(out, len(ans))
	for _, x := range ans {
		fmt.Fprint(out, x, " ")
	}

}

func pow(a, r int) int {
	res := 1
	for r > 0 {
		if r&1 == 1 {
			res *= a
		}
		a *= a
		r >>= 1
	}
	return res
}

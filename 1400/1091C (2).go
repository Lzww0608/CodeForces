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

	var n int
	fmt.Fscan(in, &n)

	calc := func(x int) int {
		return n / x * (n - x + 2) / 2
	}

	ans := []int{}
	for i := 1; i*i <= n; i++ {
		if n%i != 0 {
			continue
		}
		ans = append(ans, calc(i))
		if i*i != n {
			ans = append(ans, calc(n/i))
		}
	}

	sort.Ints(ans)
	for _, x := range ans {
		fmt.Fprint(out, x, " ")
	}
}

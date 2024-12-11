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
	ans := [][2]int{}
	for i := 0; i < n; i++ {
		ans = append(ans, [2]int{i + 1, (i+1)%n + 1})
	}
	for i := 0; i < n && !check(len(ans)); i++ {
		ans = append(ans, [2]int{i + 1, i + n/2 + 1})
	}
	fmt.Fprintln(out, len(ans))
	for _, v := range ans {
		fmt.Fprintln(out, v[0], v[1])
	}
}

func check(x int) bool {
	for i := 2; i*i <= x; i++ {
		if x%i == 0 {
			return false
		}
	}
	return true
}

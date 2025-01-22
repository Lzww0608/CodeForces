package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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
	var n int
	fmt.Fscan(in, &n)
	s := strconv.Itoa(n)
	ans := [][]int{}
	for a := 1; a <= 10_000; a++ {
		t := s
		m := a * len(s)
		for len(t) < 7 {
			t += s
		}

		for b := max(1, m-7); b <= min(m-1, 10_000); b++ {
			tmp := strconv.Itoa(n*a - b)
			if tmp == t[:m-b] {
				ans = append(ans, []int{a, b})
			}
		}
	}
	fmt.Fprintln(out, len(ans))
	for _, v := range ans {
		fmt.Fprintln(out, v[0], v[1])
	}
}

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

	var t int
	for fmt.Fscan(in, &t); t > 0; t-- {
		solve(in, out)
	}
}

func solve(in *bufio.Reader, out *bufio.Writer) {
	var n int
	fmt.Fscan(in, &n)
	ans := [][]int{}
	for n > 2 {
		s := int(math.Sqrt(float64(n)))
		if s*s < n {
			s++
		}
		for i := s + 1; i < n; i++ {
			ans = append(ans, []int{i, i + 1})
		}
		ans = append(ans, []int{n, s})
		ans = append(ans, []int{n, s})
		n = s
	}
	fmt.Fprintln(out, len(ans))
	for _, v := range ans {
		fmt.Fprintln(out, v[0], v[1])
	}
}

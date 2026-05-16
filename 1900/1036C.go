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
	var L, R int
	fmt.Fscan(in, &L, &R)

	fmt.Fprintln(out, calc(R)-calc(L-1))
}

func calc(x int) int {
	if x == 0 {
		return 0
	}
	s := strconv.Itoa(x)
	n := len(s)
	memo := make(map[[2]int]int)
	var dfs func(int, int, bool) int
	dfs = func(i, cnt int, isLimit bool) int {
		if cnt > 3 {
			return 0
		} else if i == n {
			if cnt == 0 {
				return 0
			}
			return 1
		}

		if !isLimit {
			if v, ok := memo[[2]int{i, cnt}]; ok {
				return v
			}
		}

		l, r := 0, 9
		if isLimit {
			r = int(s[i] - '0')
		}

		res := 0
		for d := l; d <= r; d++ {
			cur := cnt
			if d != 0 {
				cur++
			}

			res += dfs(i+1, cur, isLimit && d == r)
		}

		if !isLimit {
			memo[[2]int{i, cnt}] = res
		}
		return res
	}

	return dfs(0, 0, true)
}

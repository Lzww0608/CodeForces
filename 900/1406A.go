package main

import (
	"bufio"
	"fmt"
	"os"
)

const N int = 101

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
	var n, x int
	fmt.Fscan(in, &n)
	m := make(map[int]int)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &x)
		m[x]++
	}
	ans := 0
	for i := 0; i <= N; i++ {
		if m[i] < 2 {
			ans += i
			break
		}
	}

	for i := 0; i <= N; i++ {
		if m[i] < 1 {
			ans += i
			break
		}
	}

	fmt.Fprintln(out, ans)
}

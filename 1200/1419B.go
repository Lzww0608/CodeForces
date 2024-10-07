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

	var t int
	for fmt.Fscan(in, &t); t > 0; t-- {
		solve(in, out)
	}
}

func solve(in *bufio.Reader, out *bufio.Writer) {
	var n int
	fmt.Fscan(in, &n)
	ans, sum := 0, 0
	pre := 0
	for i := 1; sum+pre*2+i*i <= n; i *= 2 {
		sum, pre = sum+pre*2+i*i, pre*2+i*i
		ans++
	}
	fmt.Fprintln(out, ans)

}

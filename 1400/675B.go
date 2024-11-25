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

	var n, a, b, c, d int
	ans := 0
	fmt.Fscan(in, &n, &a, &b, &c, &d)
	for i := 1; i <= n; i++ {
		s := a + b + i
		if !(s <= a+c || s <= c+d || s <= b+d || s > a+c+n || s > c+d+n || s > b+d+n) {
			ans += n
		}
	}
	fmt.Fprintln(out, ans)
}

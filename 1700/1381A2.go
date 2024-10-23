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
	var ss, t string
	fmt.Fscan(in, &ss, &t)
	s := []byte(ss)
	ans := make([]int, 0, n<<1)
	for i := 0; i < n; i++ {
		j := i / 2
		if i&1 == 1 {
			j = n - j - 1
		}
		flip := i & 1
		x, y := int(s[j]-'0'), int(t[n-i-1]-'0')
		if x^flip == y {
			//s[j] = byte('0' + x ^ 1)
			ans = append(ans, 1)
		}
		ans = append(ans, n-i)
	}
	fmt.Fprint(out, len(ans), " ")
	for _, x := range ans {
		fmt.Fprint(out, x, " ")
	}
	fmt.Fprintln(out)
}

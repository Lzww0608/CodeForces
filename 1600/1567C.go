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
	cnt := [2]int{}
	s := strconv.Itoa(n)
	for i := range s {
		n := int(s[i] - '0')
		cnt[i%2] = cnt[i%2]*10 + n%10
		n /= 10
	}

	fmt.Fprintln(out, (cnt[0]+1)*(cnt[1]+1)-2)
}

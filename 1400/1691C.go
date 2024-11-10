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
	var n, k int
	fmt.Fscan(in, &n, &k)
	s, _ := in.ReadString('\n')
	s, _ = in.ReadString('\n')
	s = s[:len(s)-1]
	//fmt.Fprintln(out, s)

	sum := 0
	for i := 0; i < n; i++ {
		if s[i] == '1' {
			if i+1 < n {
				sum += 10
			}
			if i > 0 {
				sum += 1
			}
		}
	}

	pos := n - 1
	if s[n-1] == '0' {
		t := k
		for pos >= 0 && s[pos] == '0' && t > 0 {
			t--
			pos--
		}
		if pos >= 0 && s[pos] == '1' && t >= 0 {
			sum -= 10
			if pos == 0 {
				sum += 1
			}
			k = t
		}
	}

	if s[0] == '0' && k > 0 {
		i := 0
		for s[i] == '0' && i < pos && k > 0 {
			k--
			i++
		}
		if k >= 0 && i < pos && s[i] == '1' {
			sum -= 1
		}
	}

	fmt.Fprintln(out, sum)

}

package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)
	var s string
	fmt.Fscan(in, &s)
	ans := []byte(s)
	for i := 0; i < n; i++ {
		tmp := make([]byte, n)
		x := (10 - int(s[i]-'0')) % 10
		for j := range tmp {
			tmp[(j-i+n)%n] = byte((int(s[j]-'0')+x)%10 + '0')
		}
		if bytes.Compare(tmp, ans) < 0 {
			ans = tmp
		}
	}
	fmt.Fprintln(out, string(ans))
}

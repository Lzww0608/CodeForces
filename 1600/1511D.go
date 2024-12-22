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

	var n, k int
	fmt.Fscan(in, &n, &k)

	ans := []byte{}
	for i := 0; i < k; i++ {
		ans = append(ans, byte('a'+i))
		for j := i + 1; j < k; j++ {
			ans = append(ans, byte('a'+i))
			ans = append(ans, byte('a'+j))
		}
	}
	for n >= len(ans) {
		n -= len(ans)
		fmt.Fprint(out, string(ans))
	}
	fmt.Fprintln(out, string(ans[:n]))
}

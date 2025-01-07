package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)
	if n == 0 {
		fmt.Fprintln(out, "a")
		return
	}
	var ans string
	for i := 0; i < 26 && n > 0; i++ {
		cnt := 2
		for cnt*(cnt-1)/2 <= n {
			cnt++
		}
		cnt--
		n -= cnt * (cnt - 1) / 2
		ans += strings.Repeat(string('a'+i), cnt)
	}
	fmt.Fprintln(out, ans)
}

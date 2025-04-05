package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)
	var s string
	fmt.Fscan(in, &s)
	cnt := strings.Count(s, "?")
	if cnt&1 == 1 {
		fmt.Fprintln(out, "Monocarp")
	}

	var ans float64 = 0
	for i := 0; i < n; i++ {
		if s[i] != '?' {
			if i < n/2 {
				ans += float64(s[i] - '0')
			} else {
				ans -= float64(s[i] - '0')
			}
		} else {
			if i < n/2 {
				ans += 4.5
			} else {
				ans -= 4.5
			}
		}
	}

	if math.Abs(ans) > 1e-5 {
		fmt.Fprintln(out, "Monocarp")
	} else {
		fmt.Fprintln(out, "Bicarp")
	}
}

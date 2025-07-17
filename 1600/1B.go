package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)
	var s string
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &s)
		pos := strings.Index(s, "C")
		if pos == -1 || pos == 0 || !unicode.IsDigit(rune(s[pos-1])) {
			j := 0
			r := 0
			for s[j] >= 'A' && s[j] <= 'Z' {
				r = r*26 + int(s[j]-'A'+1)
				j++
			}
			fmt.Fprintf(out, "R%sC%d\n", s[j:], r)
		} else {
			x, _ := strconv.Atoi(s[pos+1:])
			ans := []byte{}
			for x > 0 {
				y := x % 26
				if y == 0 {
					x--
					y = 26
				}
				ans = append(ans, byte('A'+y-1))
				x = x / 26
			}
			slices.Reverse(ans)
			fmt.Fprintln(out, string(ans)+s[1:pos])
		}
	}
}

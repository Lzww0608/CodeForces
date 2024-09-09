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

	var t, n int
	for fmt.Fscan(in, &t); t > 0; t-- {
		fmt.Fscan(in, &n)
		if (n/2)%10 != 9 {
			fmt.Fprintln(out, n/2, n-n/2)
		} else {
			var a, b string
			s := strconv.Itoa(n)
			for len(s) > 0 {
				tmp := int(s[len(s)-1] - '0')
				x, y := tmp/2, (tmp+1)/2
				if a > b {
					a = string('0'+x) + a
					b = string('0'+y) + b
				} else {
					b = string('0'+x) + b
					a = string('0'+y) + a
				}
				s = s[:len(s)-1]
			}
			x, _ := strconv.Atoi(a)
			y, _ := strconv.Atoi(b)
			fmt.Fprintln(out, x, y)
		}
	}
}

package main

import (
	"bufio"
	"fmt"
	"os"
)

const N int = 101

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)
	for i := 0; i < n; i++ {
		var s string
		fmt.Fscan(in, &s)
		sum, zero, even := 0, 0, 0
		for i := range s {
			x := int(s[i] - '0')
			sum += x
			if x == 0 {
				zero++
			}
			if x&1 == 0 {
				even++
			}
		}
		if zero > 0 && even > 1 && sum%3 == 0 {
			fmt.Fprintln(out, "red")
		} else {
			fmt.Fprintln(out, "cyan")
		}
	}
}

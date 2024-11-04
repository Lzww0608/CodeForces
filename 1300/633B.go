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

	var n int
	fmt.Fscan(in, &n)
	ans := 0
	sum := 0
	for sum < n {
		ans += 5
		for t := 5; ans%t == 0; t *= 5 {
			sum++
		}
	}
	if sum > n {
		fmt.Fprintln(out, 0)
		return
	}
	fmt.Fprintln(out, 5)
	for i := ans; i < ans+5; i++ {
		fmt.Fprint(out, i, " ")
	}
}

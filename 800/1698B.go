package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var t, n, k int
	var a [200005]int

	fmt.Fscan(reader, &t)
	for t > 0 {
		t--
		fmt.Fscan(reader, &n, &k)

		for i := 0; i < n; i++ {
			fmt.Fscan(reader, &a[i])
		}

		ans := 0
		if k == 1 {
			ans = (n - 1) / 2
		} else {
			for i := 1; i < n-1; i++ {
				if a[i] > a[i-1]+a[i+1] {
					ans++
				}
			}
		}

		fmt.Fprintln(writer, ans)
	}
}

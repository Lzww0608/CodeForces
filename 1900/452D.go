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

	var k, n1, n2, n3, t1, t2, t3 int
	fmt.Fscan(in, &k, &n1, &n2, &n3, &t1, &t2, &t3)

	ans := 0
	f1 := make([]int, n1)
	f2 := make([]int, n2)
	f3 := make([]int, n3)

	for i := 0; i < k; i++ {
		ans = max(f1[i%n1]+t1+t2+t3, f2[i%n2]+t2+t3, f3[i%n3]+t3)
		f1[i%n1] = ans - t2 - t3
		f2[i%n2] = ans - t3
		f3[i%n3] = ans
	}

	fmt.Fprintln(out, ans)
}

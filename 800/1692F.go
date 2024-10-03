package main

import (
	"bufio"
	. "fmt"
	"os"
)

func main() {

	var t int
	in := bufio.NewReader(os.Stdin)
	Fscan(in, &t)
outer:
	for ; t > 0; t-- {

		var n, x int
		Fscan(in, &n)
		a := [10]int{}
		for i := 0; i < n; i++ {
			Fscan(in, &x)
			a[x%10]++
		}
		for i := 0; i < 10; i++ {
			if a[i] == 0 {
				continue
			}
			a[i]--
			for j := i; j < 10; j++ {
				if a[j] == 0 {
					continue
				}
				a[j]--
				for k := j; k < 10; k++ {
					if a[k] > 0 && (i+j+k)%10 == 3 {
						Println("YES")
						continue outer
					}
				}
				a[j]++
			}
			a[i]++
		}
		Println("NO")

	}

	return
}

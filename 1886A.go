package main

import (
	"bufio"
	. "fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	var t, n int

next:
	for Fscan(in, &t); t > 0; t-- {
		Fscan(in, &n)
		for a := 1; a < 10; a++ {
			for b := a + 1; b < 10; b++ {
				if a%3 != 0 && b%3 != 0 && (n-a-b)%3 != 0 && a != n-a-b && b != n-a-b && n-a-b > 0 {
					Println("YES")
					Println(a, b, n-a-b)
					continue next
				}
			}
		}
		Println("NO")
	}
}

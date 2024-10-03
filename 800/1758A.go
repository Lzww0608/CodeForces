package main

import (
	"bufio"
	. "fmt"
	"os"
)

func main() {
	in := bufio.NewScanner(os.Stdin)
	var t int
	for Scanf("%d\n", &t); t > 0; t-- {
		in.Scan()
		s := string(in.Text())
		Print(s)
		for n := len(s) - 1; n >= 0; n-- {
			Print(string(s[n]))
		}
		Println()

	}
}

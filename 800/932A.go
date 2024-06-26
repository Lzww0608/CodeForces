package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)
	for _, c := range s {
		fmt.Print(string(c))
	}
	for i := len(s) - 1; i >= 0; i-- {
		fmt.Print(string(s[i]))
	}
	fmt.Println()
}

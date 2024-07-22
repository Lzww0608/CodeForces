package main

import "fmt"

func main() {
	var s, t string
	fmt.Scan(&s)
	fmt.Scan(&t)
	if len(s) != len(t) {
		fmt.Println(max(len(s), len(t)))
		return
	}

	if s != t {
		fmt.Println(len(s))
	} else {
		fmt.Println(-1)
	}
}

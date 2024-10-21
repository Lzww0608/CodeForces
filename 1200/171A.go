package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	s := []byte(strconv.Itoa(b))
	n := len(s)
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	t, _ := strconv.Atoi(string(s))
	fmt.Println(t + a)
}

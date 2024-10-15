package main

import (
	"fmt"
	"sort"
)

func main() {
	a := make([]int, 3)
	fmt.Scan(&a[0], &a[1], &a[2])
	sort.Ints(a)
	if a[0] == a[1] && a[1] == a[2] {
		fmt.Println(0)
	} else if a[1] == a[2] {
		fmt.Println(a[1] - a[0] - 1)
	} else if a[0] == a[1] {
		fmt.Println(2*(a[2]-a[1]) - 2)
	} else {
		fmt.Println(a[2]*2 - a[1] - a[0] - 2)
	}
}

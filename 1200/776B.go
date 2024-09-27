package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	if n == 2 || n == 1 {
		fmt.Println(1)
		for i := 2; i < n+2; i++ {
			fmt.Print(1, " ")
		}
		fmt.Println()
		return
	}
	primes := make([]int, n+2)
	a := make([]int, n+2)
	for i := 2; i*i <= len(primes); i++ {
		if a[i] == 0 {
			a[i] = 1
		}
		if primes[i] == 1 {
			continue
		}
		for j := i * i; j < len(primes); j += i {
			primes[j] = 1
			a[j] = 2
		}
	}
	fmt.Println(2)
	for i := 2; i < len(a); i++ {
		fmt.Print(max(1, a[i]), " ")
	}
	fmt.Println()
}

package main

import "fmt"

func main() {
	var n, a, b int
	fmt.Scan(&n)
	fmt.Scan(&a)
	fmt.Scan(&b)
	if n == 1 {
		if a == b {
			fmt.Println(1)
		} else {
			fmt.Println(0)
		}
		fmt.Printf("%d:%d\n", a, b)
		return
	}

	if a+b <= n {
		fmt.Println(n - a - b)
		for i := 0; i < n; i++ {
			if a > 0 {
				a--
				fmt.Println("1:0")
			} else if b > 0 {
				b--
				fmt.Println("0:1")
			} else {
				fmt.Println("0:0")
			}
		}
		return
	}

	if a+b > n {
		if a == 0 {
			fmt.Println(0)
			for i := 0; i < n; i++ {
				if i == n-1 {
					fmt.Printf("0:%d\n", b)
				} else {
					b--
					fmt.Println("0:1")
				}
			}
		} else if b == 0 {
			fmt.Println(0)
			for i := 0; i < n; i++ {
				if i == n-1 {
					fmt.Printf("%d:0\n", a)
				} else {
					a--
					fmt.Println("1:0")
				}
			}
		} else {
			fmt.Println(0)
			for i := 0; i < n; i++ {
				if i+2 == n {
					fmt.Printf("%d:0\n", a)
					fmt.Printf("0:%d\n", b)
					return
				} else if a > 1 {
					a--
					fmt.Println("1:0")
				} else if b > 1 {
					b--
					fmt.Println("0:1")
				}
			}
		}
	}

}

package main

import "fmt"

func main() {
	var d, sum int

	fmt.Scan(&d, &sum)
	a := make([][2]int, d)
	minSum, maxSum := 0, 0
	for i := range a {
		fmt.Scan(&a[i][0], &a[i][1])
		minSum += a[i][0]
		maxSum += a[i][1]
	}
	if sum < minSum || sum > maxSum {
		fmt.Println("NO")
		return
	}
	fmt.Println("YES")
	for i := 0; i < d; i++ {
		x, y := a[i][0], a[i][1]
		minSum -= x
		maxSum -= y
		if sum-x >= minSum && sum-x <= maxSum {
			sum -= x
			fmt.Print(x, " ")
		} else {
			fmt.Print(min(y, sum), " ")
			sum -= min(y, sum)
		}
	}
	fmt.Println()
}

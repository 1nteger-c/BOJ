package main

import "fmt"

func main() {
	var n, x, y int
	fmt.Scanf("%d\n", &n)
	for i := 1; i <= n; i++ {
		fmt.Scanf("%d %d\n", &x, &y)
		fmt.Printf("Case #%d: %d\n", i, x+y)
	}
}

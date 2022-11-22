package main

import "fmt"

func main() {
	var n, x int
	fmt.Scanf("%d %d\n", &n, &x)
	for i := 0; i < n; i++ {
		data := 0
		fmt.Scanf("%d", &data)
		if data < x {
			fmt.Printf("%d ", data)
		}
	}
}

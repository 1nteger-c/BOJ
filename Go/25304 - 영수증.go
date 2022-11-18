package main

import "fmt"

func main() {
	var X, N, a, b int
	fmt.Scanf("%d\n%d\n", &X, &N)
	for i := 0; i < N; i++ {
		fmt.Scanf("%d %d\n", &a, &b)
		X -= a * b
	}
	if X == 0 {
		fmt.Printf("Yes")
	} else {
		fmt.Printf("No")
	}
}

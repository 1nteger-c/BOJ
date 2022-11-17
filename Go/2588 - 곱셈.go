package main

import "fmt"

func main() {
	var a, b int
	fmt.Scanf("%d\n%d", &a, &b)
	var c = a * b
	for i := 0; i < 3; i++ {
		n := b % 10
		b = b / 10
		fmt.Printf("%d\n", a*n)
	}
	fmt.Printf("%d", c)
}

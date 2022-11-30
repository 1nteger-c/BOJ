package main

import "fmt"

func main() {
	var n int
	sum := 0
	fmt.Scanf("%d\n", &n)
	for i := 0; i < n; i++ {
		c := 0
		fmt.Scanf("%c", &c)
		c = c - '0'
		sum += c
	}
	fmt.Printf("%d\n", sum)

}

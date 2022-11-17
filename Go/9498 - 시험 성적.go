package main

import "fmt"

func main() {
	var num int
	fmt.Scanf("%d", &num)
	if num >= 90 {
		fmt.Printf("A")
	} else if num >= 80 {
		fmt.Printf("B")
	} else if num >= 70 {
		fmt.Printf("C")
	} else if num >= 60 {
		fmt.Printf("D")
	} else {
		fmt.Printf("F")
	}
}

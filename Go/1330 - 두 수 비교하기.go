package main

import "fmt"

func main() {
	var num1, num2 int
	fmt.Scanf("%d %d", &num1, &num2)
	if num1 < num2 {
		fmt.Printf("<")
	} else if num1 == num2 {
		fmt.Printf("==")
	} else {
		fmt.Printf(">")
	}
}

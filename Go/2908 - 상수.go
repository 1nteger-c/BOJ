package main

import "fmt"

func main() {
	n1, n2 := 0, 0
	fmt.Scanf("%d %d\n", &n1, &n2)
	num1, num2 := 0, 0
	for i := 0; i < 3; i++ {
		num1 *= 10
		num2 *= 10
		num1 += n1 % 10
		num2 += n2 % 10
		n1 /= 10
		n2 /= 10
	}
	if num1 > num2 {
		fmt.Printf("%d\n", num1)
	} else {
		fmt.Printf("%d\n", num2)
	}

}

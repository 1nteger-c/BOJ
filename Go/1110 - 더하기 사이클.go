package main

import "fmt"

func main() {
	var n, x, cnt int
	cnt = 0
	fmt.Scanf("%d", &n)
	x = n
	for true {
		cnt += 1
		a, b := x/10, x%10
		c := a + b
		x = 10*b + c%10
		if x == n {
			break
		}
	}
	fmt.Printf("%d", cnt)
}

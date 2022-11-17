package main

import "fmt"

func main() {
	var hour, min, c int
	fmt.Scanf("%d %d\n%d", &hour, &min, &c)
	min += c
	x := min / 60
	min = min % 60
	hour = (hour + x) % 24
	fmt.Printf("%d %d", hour, min)
}

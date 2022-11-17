package main

import "fmt"

func main() {
	var hour, min int
	fmt.Scanf("%d %d", &hour, &min)
	min -= 45
	if min < 0 {
		min += 60
		hour -= 1
		if hour < 0 {
			hour += 24
		}
	}
	fmt.Printf("%d %d", hour, min)
}

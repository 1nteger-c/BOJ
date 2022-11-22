package main

import "fmt"

func main() {
	var arr = make([]int, 31)
	var inp int
	for i := 0; i < 28; i++ {
		fmt.Scanf("%d\n", &inp)
		arr[inp] = 1
	}
	for i := 1; i <= 30; i++ {
		if arr[i] == 0 {
			fmt.Printf("%d\n", i)
		}
	}
}

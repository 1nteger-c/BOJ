package main

import "fmt"

func main() {
	var num [6]int
	var cnt = [6]int{1, 1, 2, 2, 2, 8}
	for i := 0; i < 6; i++ {
		fmt.Scanf("%d", &num[i])
	}
	for i := 0; i < 6; i++ {
		fmt.Printf("%d ", cnt[i]-num[i])
	}

}

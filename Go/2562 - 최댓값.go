package main

import "fmt"

func main() {
	var data, idx int
	var max int = 0
	for i := 1; i < 10; i++ {
		fmt.Scanf("%d\n", &data)
		if max < data {
			max = data
			idx = i
		}
	}
	fmt.Printf("%d\n%d", max, idx)
}

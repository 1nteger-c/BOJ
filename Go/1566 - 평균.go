package main

import "fmt"

func main() {
	var res float32
	var n int
	var max int = 0
	fmt.Scanf("%d\n", &n)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &arr[i])
		if max < arr[i] {
			max = arr[i]
		}
	}
	for i := 0; i < n; i++ {
		res += ((float32(arr[i]) / float32(max)) * 100)
	}
	fmt.Printf("%v", res/float32(n))
}

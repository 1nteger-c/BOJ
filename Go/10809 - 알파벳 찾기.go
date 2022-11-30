package main

import "fmt"

func main() {
	var arr = make([]byte, 100)
	fmt.Scanf("%s\n", &arr)
	var ptr = make([]int, 27)
	for i := 0; i < len(arr); i++ {
		x := arr[i] - 'a'
		if ptr[x] == 0 {
			ptr[x] = i + 1
		}
	}
	for i := 0; i < 26; i++ {
		fmt.Printf("%d ", ptr[i]-1)
	}
}

package main

import "fmt"

func main() {
	arr := make([]int, 42)
	cnt := 0
	var data int
	for i := 0; i < 10; i++ {
		fmt.Scanf("%d\n", &data)
		x := data % 42
		if arr[x] == 0 {
			cnt += 1
			arr[x] = 1
		}
	}
	fmt.Printf("%d\n", cnt)
}

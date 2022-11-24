package main

import "fmt"

func d(x int) int {
	ret := x
	for x != 0 {
		ret += x % 10
		x = x / 10
	}
	return ret
}

func main() {
	var array = make([]int, 10001)
	for i := 1; i < 10001; i++ {
		a := d(i)
		if a < 10001 {
			array[a] = 1
		}
		if array[i] == 0 {
			fmt.Printf("%d\n", i)
		}
	}

}

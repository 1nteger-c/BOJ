package main

import "fmt"

func main() {
	var n int
	var x = [100]int{}
	fmt.Scanf("%d\n", &n)
	for i := 0; i < n; i++ {
		if i < n-1 {
			fmt.Scanf("%d", &x[i])
		} else {
			fmt.Scanf("%d\n", &x[i])
		}
	}
	var goal int
	var cnt int = 0
	fmt.Scanf("%d", &goal)
	for i := 0; i < n; i++ {
		if x[i] == goal {
			cnt += 1
		}
	}
	fmt.Printf("%d", cnt)
}

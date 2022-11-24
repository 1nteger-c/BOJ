package main

import "fmt"

func main() {
	var n int
	var cnt int = 0
	fmt.Scanf("%d", &n)
	for i := 1; i < n+1; i++ {
		if i < 100 {
			cnt += 1
			continue
		}
		x := i
		a0 := x % 10
		x = x / 10
		a1 := x % 10
		x = x / 10
		flag := 0
		for x > 0 {
			if x%10-a1 != a1-a0 {
				flag = 1
				break
			}
			a0 = a1
			a1 = x % 10
			x = x / 10
		}
		if flag == 0 {
			cnt += 1

		}
	}
	fmt.Printf("%d\n", cnt)

}

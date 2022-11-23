package main

func sum(a []int) int {
	var x int = 0
	for i := 0; i < len(a); i++ {
		x += a[i]
	}
	return x
}

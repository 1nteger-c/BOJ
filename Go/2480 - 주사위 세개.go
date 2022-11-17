package main

import "fmt"

func main() {
	var x, y, z int
	fmt.Scanf("%d %d %d", &x, &y, &z)
	if x < y {
		x, y = y, x
	}
	if y < z {
		y, z = z, y
	}
	if x < y {
		x, y = y, x
	}
	if x == y && y == z {
		fmt.Printf("%d", 10000+x*1000)
	} else if x == y {
		fmt.Printf("%d", 1000+x*100)
	} else if y == z {
		fmt.Printf("%d", 1000+y*100)
	} else {
		fmt.Printf("%d", x*100)
	}

}

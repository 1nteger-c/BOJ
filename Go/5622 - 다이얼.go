package main

import "fmt"

func main() {
	in := make([]byte, 15)
	fmt.Scanf("%s\n", &in)
	sum := 0
	for i := 0; i < len(in); i++ {
		x := in[i]
		if x == 'Z' {
			x -= 1
		}
		if x >= 'S' {
			x -= 1
		}
		sum += int(3 + (x-'A')/3)

	}
	fmt.Printf("%d\n", sum)
}

package main

import "fmt"

func main() {
	var n int
	fmt.Scanf("%d\n", &n)
	for i := 0; i < n; i++ {
		l := 0
		fmt.Scanf("%d", &l)
		x := make([]byte, 20)
		fmt.Scanf("%s\n", &x)
		for j := 0; j < len(x); j++ {
			for k := 0; k < l; k++ {
				fmt.Printf("%c", x[j])
			}
		}
		fmt.Print("\n")
	}
}

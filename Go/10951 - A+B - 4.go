package main

import (
	"fmt"
	"io"
)

func main() {
	var a, b int
	for true {
		_, err := fmt.Scanf("%d %d", &a, &b)
		if err == io.EOF {
			break
		}
		fmt.Printf("%d\n", a+b)
	}
}

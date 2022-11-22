package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	var n int
	fmt.Scanf("%d\n", &n)
	for i := 0; i < n; i++ {
		bytes, _, _ := r.ReadLine()
		line := string(bytes)
		score := 0
		pt := 1
		for j := 0; j < len(line); j++ {
			if line[j] == 'O' {
				score += pt
				pt += 1
			} else {
				pt = 1
			}
		}
		fmt.Printf("%d\n", score)
	}
}

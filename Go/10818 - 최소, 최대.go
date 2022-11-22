package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()
	min := 1000001
	max := -1000001

	var n, data int
	fmt.Scanf("%d\n", &n)
	for i := 0; i < n; i++ {
		fmt.Fscan(r, &data)
		if data < min {
			min = data
		}
		if data > max {
			max = data
		}
	}
	fmt.Printf("%d %d", min, max)

}

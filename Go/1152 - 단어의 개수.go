package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	r := bufio.NewReaderSize(os.Stdin, 1000000)
	arr, _, _ := r.ReadLine()
	line := string(arr)
	nums := strings.Fields(line)
	fmt.Printf("%d\n", len(nums))
}

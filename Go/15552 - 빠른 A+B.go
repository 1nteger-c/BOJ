package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	bytes, _, _ := r.ReadLine()
	line := string(bytes)
	n, _ := strconv.Atoi(line)
	for i := 0; i < n; i++ {
		bytes, _, _ := r.ReadLine()
		line := string(bytes)
		nums := strings.Fields(line)
		a, _ := strconv.Atoi(nums[0])
		b, _ := strconv.Atoi(nums[1])
		w.WriteString(strconv.Itoa(a+b) + "\n")
	}
}

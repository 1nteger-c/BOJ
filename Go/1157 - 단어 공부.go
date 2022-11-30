package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	arr := make([]byte, 1000000)
	r := bufio.NewReader(os.Stdin)
	fmt.Fscanf(r, "%s\n", &arr)
	cnt := make([]int, 26)
	for i := 0; i < len(arr); i++ {
		x := arr[i]
		if x < 'a' {
			x = x + 'a' - 'A'
		}
		cnt[x-'a'] += 1
	}
	ptr := -1
	max := -1
	dup := false
	for i := 0; i < len(cnt); i++ {
		if max < cnt[i] {
			ptr = i
			max = cnt[i]
			dup = false
		} else if max == cnt[i] {
			dup = true
		}
	}
	if dup {
		fmt.Print("?")
	} else {
		fmt.Printf("%c", ptr+'A')
	}
}

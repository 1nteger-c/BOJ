package main

import "fmt"

func main() {
	n := 0
	cnt := 0
	fmt.Scanf("%d\n", &n)
	for i := 0; i < n; i++ {
		in := make([]byte, 100)
		data := make([]int, 26)
		fmt.Scanf("%s\n", &in)
		st := -1
		flag := false
		for j := 0; j < len(in); j++ {
			x := int(in[j] - 'a')
			if st == x {
				continue
			}
			st = x
			if data[x] != 0 {
				flag = true
				break
			}
			data[x] = 1
		}
		if flag == false {
			cnt += 1
		}
	}
	fmt.Printf("%d\n", cnt)
}

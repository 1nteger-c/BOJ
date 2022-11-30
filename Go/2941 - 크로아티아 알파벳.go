package main

import "fmt"

func main() {
	in := make([]byte, 100)
	fmt.Scanf("%s", &in)
	cnt := 0
	ptr := 0
	for ptr < len(in) {
		if ptr+2 < len(in) && string(in[ptr:ptr+3]) == "dz=" {
			ptr += 3
		} else if ptr+1 < len(in) {
			switch string(in[ptr : ptr+2]) {
			case
				"c=",
				"c-",
				"d-",
				"lj",
				"nj",
				"s=",
				"z=":
				ptr += 2
			default:
				ptr += 1
			}
		} else {
			ptr += 1
		}
		cnt += 1
	}
	fmt.Printf("%d\n", cnt)
}

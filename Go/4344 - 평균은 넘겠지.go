package main

import "fmt"

func main() {
	var n, nn int
	fmt.Scanf("%d\n", &n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &nn)
		data := make([]int, nn)
		var mean float32 = 0
		for j := 0; j < nn; j++ {
			if j != nn-1 {
				fmt.Scanf("%d", &data[j])
			} else {
				fmt.Scanf("%d\n", &data[j])
			}
			mean += float32(data[j])
		}
		var cnt = 0
		mean = mean / float32(nn)
		for j := 0; j < nn; j++ {
			if float32(data[j]) > mean {
				cnt += 1
			}
		}
		fmt.Printf("%.3f%%\n", float32(cnt)/float32(nn)*100)
	}
}

package main

import "fmt"

func main() {
	var tc int
	fmt.Scan(&tc)
	for t := 0; t < tc; t++ {
		var n int
		fmt.Scan(&n)
		arr := make([]int, n)
		for i := 0; i < n; i++ {
			var t int
			fmt.Scan(&t)
			arr[i] = t
		}
		rs := solution(arr, n)
		fmt.Println(rs)
	}
}

func solution(arr []int, n int) (rs int) {
	for i := 0; i < n; i++ {
		rs += arr[i]
	}
	return
}

package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	inType := "file"
	var arr []int
	var n int
	switch inType {
	case "file":
		strs := strings.Fields(inStr)
		n = len(strs)
		arr = make([]int, n)
		var err error
		for i := 0; i < n; i++ {
			arr[i], err = strconv.Atoi(strs[i])
			if err != nil {
				continue
			}
		}
	case "std":
		fmt.Scan(&n)
		for i := 0; i < n; i++ {
			var t int
			fmt.Scan(&t)
			arr = append(arr, t)
		}
	}
	rs := solution(arr, n)
	fmt.Println(rs)
}

func solution(arr []int, n int) (rs int) {
	prevSum := arr[0] + arr[1] + arr[2]
	for i := 3; i < n; i++ {
		currSum := arr[i-2] + arr[i-1] + arr[i]
		if currSum > prevSum {
			rs++
		}
		prevSum = currSum
	}
	return
}

//go:embed input.txt
var inStr string

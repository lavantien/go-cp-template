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
	for i := 1; i < n; i++ {
		if arr[i] > arr[i-1] {
			rs++
		}
	}
	return
}

//go:embed input.txt
var inStr string

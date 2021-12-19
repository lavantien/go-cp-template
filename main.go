package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"strings"
)

func main() {
	test := 0
	var str string
	if test == 1 {
		str = testStr
	} else {
		str = inStr
	}
	scn := bufio.NewScanner(strings.NewReader(str))
	var wordSize int
	n := 0
	n++
	scn.Scan()
	tstr := scn.Text()
	wordSize = len(tstr)
	count := make([]int, wordSize)
	for i := 0; i < wordSize; i++ {
		if tstr[i] == '1' {
			count[i]++
		}
	}
	for scn.Scan() {
		n++
		tstr := scn.Text()
		for i := 0; i < wordSize; i++ {
			if tstr[i] == '1' {
				count[i]++
			}
		}
	}
	gamma := make([]int, wordSize)
	epsilon := make([]int, wordSize)
	halfSize := n / 2
	for i := 0; i < wordSize; i++ {
		if count[i] > halfSize {
			gamma[i] = 1
			epsilon[i] = 0
		} else {
			gamma[i] = 0
			epsilon[i] = 1
		}
	}
	rs := wordToInt(gamma[:]) * wordToInt(epsilon[:])
	fmt.Println(rs)
}

//go:embed input.txt
var inStr string

//go:embed test.txt
var testStr string

func wordToInt(word []int) (rs int) {
	n := len(word)
	for i := n - 1; i >= 0; i-- {
		rs += word[i] * pw2(n-i-1)
	}
	return
}

func pw2(x int) (rs int) {
	rs = 1
	for i := 0; i < x; i++ {
		rs *= 2
	}
	return
}

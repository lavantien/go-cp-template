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
	distance := 0
	depth := 0
	aim := 0
	for scn.Scan() {
		var cmd string
		var val int
		fmt.Sscanln(scn.Text(), &cmd, &val)
		switch cmd {
		case "forward":
			distance += val
			depth += aim * val
		case "up":
			aim -= val
		case "down":
			aim += val
		}
	}
	rs := distance * depth
	fmt.Println(rs)
}

//go:embed input.txt
var inStr string

//go:embed test.txt
var testStr string

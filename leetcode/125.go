package main

import (
	"fmt"
	"regexp"
	"strings"
)

func test() bool {
	s := "0P"

	s = strings.ToLower(s)
	compile, _ := regexp.Compile(`[^a-z]+`)
	s1 := compile.ReplaceAllString(s, "")
	fmt.Printf(s1)
	fmt.Printf("\n")
	for x, y := 0, len(s1)-1; x <= y; x++ {
		fmt.Printf("%v,%v\n", string(s1[x]), string(s1[y]))
		if string(s1[x]) != string(s1[y]) {
			return false
		}
		y--
	}
	return true
}
func main() {
	fmt.Printf("%v", test())
	//strings.Join()
}

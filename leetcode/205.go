package main

import (
	"fmt"
)

func main() {
	fmt.Printf("%v\n", isIsomorphic("egg", "add"))

}

func isIsomorphic(s string, t string) bool {
	sMap := map[byte]byte{}
	tMap := map[byte]byte{}
	var s2, t2 byte

	for i := 0; i < len(s); i++ {
		s2, t2 = s[i], t[i]

		if _, ok := sMap[s2]; ok {
			if sMap[s2] != t2 {
				return false
			}
		} else {
			sMap[s2] = t2
		}

		if _, ok := tMap[t2]; ok {
			if tMap[t2] != s2 {
				return false
			}
		} else {
			tMap[t2] = s2
		}
	}
	return true
}

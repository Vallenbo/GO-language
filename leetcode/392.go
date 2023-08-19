package main

func isSubsequence(s string, t string) bool {
	lens1 := len(s)
	lens2 := len(t)
	str := ""
	for x, y := 0, 0; x < lens1 && y < lens2; {
		if string(s[x]) != string(t[y]) {
			y++
		} else {
			str = str + string(s[x])
			y++
			x++
		}

	}

	if str != s {
		return false
	}
	return true
}

func main() {
	s := "axc"
	t := "ahbgdc"

	println(isSubsequence(s, t))

}

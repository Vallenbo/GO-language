package main

func Test_lengthOfLongestSubstring() {
	lengthOfLongestSubstring("werss")
}
func lengthOfLongestSubstring(s string) int {
	lastOccurred := make(map[byte]int)
	var start, maxLength int
	for i := 0; i < len(s); i++ {
		ch := s[i]
		if index, ok := lastOccurred[ch]; ok && index >= start {
			start = index + 1
		}
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccurred[ch] = i
	}
	return maxLength
}

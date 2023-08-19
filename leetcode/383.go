package main

func canConstruct(ransomNote string, magazine string) bool {
	charCount := make(map[rune]int) // 创建哈希表用于记录 magazine 中每个字符的出现次数

	for _, ch := range magazine { // 统计 magazine 中每个字符的出现次数
		charCount[ch]++
	}

	for _, ch := range ransomNote { // 检查 ransomNote 中的字符是否可以由 magazine 构成
		if charCount[ch] == 0 {
			return false
		}
		charCount[ch]--
	}

	return true
}

func main() {
	canConstruct("abc", "abcdefg")
}

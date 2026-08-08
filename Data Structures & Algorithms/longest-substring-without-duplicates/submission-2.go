func lengthOfLongestSubstring(s string) int {
	maxLen := 0
	charMap := make(map[byte]int,len(s))
	left := 0

	for right := 0;right < len(s);right++ {
		if ind, ok := charMap[s[right]]; ok && ind >= left {
			left = ind  + 1
		}

		charMap[s[right]] = right

		if (right - left + 1) > maxLen {
			maxLen = right - left + 1
		}
	}

	return maxLen
}

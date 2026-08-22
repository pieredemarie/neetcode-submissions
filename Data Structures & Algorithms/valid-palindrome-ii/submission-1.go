func validRange(s string, left, right int) bool {
	for left < right {
		if s[left] != s[right] {
			return false 
		}
		left++
		right--
	}

	return true
}
func validPalindrome(s string) bool {
	left,right := 0, len(s) - 1

	for left < right {
		if s[left] != s[right] {
			return validRange(s,left+1,right) || validRange(s,left,right-1)
		}
		left++
		right--
	}

	return true
}

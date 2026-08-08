func hasDuplicate(nums []int) bool {
    n := len(nums)
	numMap := make(map[int]int,n)

	for _, num := range nums {
		numMap[num]++
		if numMap[num] > 1 {
			return true
		} 
		
	}

	return false
}
